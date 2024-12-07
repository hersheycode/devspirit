package dgraph

import (
	"context"
	"crypto/tls"
	"fmt"

	"time"

	"apppathway.com/examples/prodapi/pkg/plugins/intent"
	"apppathway.com/pkg/errors"
	dgo "github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
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

	Cache intent.CacheService[any]
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

// IntentService represents a service for managing intent.
type IntentService struct {
	*DB
}

// NewIntentService returns a new instance of IntentService.
func NewIntentService(db *DB) *IntentService {
	return &IntentService{DB: db}
}

type Intent struct {
	Name string `json:"name"`
}

func (is *IntentService) Register(ctx context.Context, req intent.RegisterIntentReq) (intent.RegisterIntentRes, error) {
	i := Intent{
		Name: req.Name,
	}
	if err := i.commit(ctx, is); err != nil {
		return intent.RegisterIntentRes{}, err
	}
	return intent.RegisterIntentRes{Status: "registered intent"}, nil
}

func (i Intent) commit(ctx context.Context, is *IntentService) error {
	query := `{
		q(func: type(Intent)) @filter(eq(name, "` + i.Name + `")) { 
			intentUID as uid
		}
	}`

	res, err := upsert[Intent](ctx, is, query, &api.Mutation{
		Cond: `@if(eq(len(intentUID), 0))`, //intent does not exist
		SetNquads: []byte(`
			_:intent <name> "` + i.Name + `" . .
			_:intent <dgraph.type> "Intent" .
		`),
	})
	if err != nil {
		return fmt.Errorf("err while upserting intent: %v", err)
	}
	if len(res) != 0 {
		return errors.ConflictError(fmt.Errorf("%s exists", i.Name))
	}
	return nil
}
