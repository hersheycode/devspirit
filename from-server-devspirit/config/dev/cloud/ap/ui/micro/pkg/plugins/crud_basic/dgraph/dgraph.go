package dgraph

import (
	"apppathway.com/pkg/errors"
	"codestore.localhost/crudusrs/crud_basic/api/crudbasic"
	"context"
	"crypto/tls"
	"fmt"
	dgo "github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
	"time"
)

type DB struct {
	*dgo.Dgraph
	ctx    context.Context
	cancel func()
	dial   *grpc.ClientConn
	DSN    string
	Now    func() time.Time
	TLS    *tls.Config
	Cache  intent.CacheService[any]
}
type Intent struct {
	Name string `json:"name"`
}
type IntentService struct{ *DB }

func (i Intent) commit(ctx context.Context, is *IntentService) error {
	query := `{
		q(func: type(Intent)) @filter(eq(name, "` + i.Name + `")) { 
			intentUID as uid
		}
	}`
	res, err := upsert[Intent](ctx, is, query, &api.Mutation{Cond: `@if(eq(len(intentUID), 0))`, SetNquads: []byte(`
			_:intent <name> "` + i.Name + `" . .
			_:intent <dgraph.type> "Intent" .
		`)})
	if err != nil {
		return fmt.Errorf("err while upserting intent: %v", err)
	}
	if len(res) != 0 {
		return errors.ConflictError(fmt.Errorf("%s exists", i.Name))
	}
	return nil
}
func New(dsn string, conf *tls.Config) *DB {
	db := &DB{DSN: dsn, Now: time.Now, TLS: conf}
	db.ctx, db.cancel = context.WithCancel(context.Background())
	return db
}
func (d *DB) Open() error {
	var err error
	if d.DSN == "" {
		return fmt.Errorf("dsn required")
	}
	dialOpts := []grpc.DialOption{}
	if d.TLS != nil {
		creds := grpc.WithTransportCredentials(credentials.NewTLS(d.TLS))
		dialOpts = append([]grpc.DialOption{}, creds, grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	} else {
		dialOpts = append([]grpc.DialOption{}, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	}
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
func NewIntentService(db *DB) *IntentService {
	return &IntentService{DB: db}
}
func (is *IntentService) Register(ctx context.Context, req intent.RegisterIntentReq) (intent.RegisterIntentRes, error) {
	i := Intent{Name: req.Name}
	if err := i.commit(ctx, is); err != nil {
		return intent.RegisterIntentRes{}, err
	}
	return intent.RegisterIntentRes{Status: "registered intent"}, nil
}
