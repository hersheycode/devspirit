package dgraph

import (
	"apppathway.com/pkg/user/behavior"
	"apppathway.com/pkg/user/behavior/recomodelcli"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/gob"
	"encoding/json"
	"fmt"
	dgo "github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
	"io"
	"time"
)

type DB struct {
	*dgo.Dgraph
	ctx    context.Context // background context
	cancel func()          // cancel background context
	dial   *grpc.ClientConn
	// Datasource name.
	DSN string
	// Returns the current time. Defaults to time.Now().
	// Can be mocked for tests.
	Now func() time.Time
	TLS *tls.Config

	Cache behavior.CacheService[any]

	RecoModelCli *recomodelcli.RecoModelCliService
}

// NewDB returns a new instance of DB associated with the given datasource name.
func New(dsn string, conf *tls.Config) *DB {
	db := &DB{
		DSN: dsn,
		Now: time.Now,
		TLS: conf,
	}
	db.ctx, db.cancel = context.WithCancel(context.Background())
	return db
}

// Open opens the database connection.
func (d *DB) Open() error {
	var err error
	// Ensure a DSN is set before attempting to open the database.
	if d.DSN == "" {
		return fmt.Errorf("dsn required")
	}
	creds := grpc.WithTransportCredentials(credentials.NewTLS(d.TLS))
	dialOpts := append([]grpc.DialOption{},
		creds,
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))

	d.dial, err = grpc.Dial(d.DSN, dialOpts...)
	if err != nil {
		return fmt.Errorf("while grpc dialing: %v", err)
	}
	d.Dgraph = dgo.NewDgraphClient(api.NewDgraphClient(d.dial))
	return nil
}

func (d *DB) Close() {
	d.dial.Close()
}

// BehaviorService represents a service for managing behavior.
type BehaviorService struct {
	*DB
}

// NewBehaviorService returns a new instance of BehaviorService.
func NewBehaviorService(db *DB) *BehaviorService {
	return &BehaviorService{DB: db}
}

func (b *BehaviorService) LogCmd(ctx context.Context, rw io.ReadWriter) error {
	serve := decodeEncode[behavior.LogCmdReq, behavior.LogCmdRes]
	return serve(rw, func(req behavior.LogCmdReq, res *behavior.LogCmdRes) error {
		err := b.RecoModelCli.Recommendation(req.Command)
		if err != nil {
			return err
		}
		res.Status = fmt.Sprintf("success: %v", req.Command)
		return nil
	})
}

func (b *BehaviorService) Temp(ctx context.Context, rw io.ReadWriter) error {
	serve := decodeEncode[behavior.LogCmdReq, behavior.LogCmdRes]
	return serve(rw, func(req behavior.LogCmdReq, res *behavior.LogCmdRes) error {
		mu := &api.Mutation{CommitNow: true}
		buf := &bytes.Buffer{}
		t := struct {
			Data string `json:"data"`
		}{
			Data: "Hello, world Req",
		}
		err := json.NewEncoder(buf).Encode(t)
		if err != nil {
			return err
		}
		mu.SetJson = buf.Bytes()
		response, err := b.NewTxn().Mutate(ctx, mu)
		if err != nil {
			return err
		}
		fmt.Printf("Database: %+v, Req: %#v \n", response, req)
		res.Status = "test successful"
		return nil
	})
}

func decodeEncode[r behavior.ReqReader, w behavior.ResWriter](rw io.ReadWriter, handler behavior.HandleFunc[r, w]) error {

	var req r
	if err := gob.NewDecoder(rw).Decode(&req); err != nil {
		return nil
	}

	var res w
	if err := handler(req, &res); err != nil {
		return err
	}

	if err := gob.NewEncoder(rw).Encode(res); err != nil {
		return err
	}
	return nil
}
