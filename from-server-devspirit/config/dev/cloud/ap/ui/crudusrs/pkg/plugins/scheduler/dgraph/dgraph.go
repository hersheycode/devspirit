package dgraph

import (
	"context"
	"crypto/tls"
	"fmt"

	"time"

	"apppathway.com/examples/prodapi/pkg/plugins/scheduler"
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

	Cache scheduler.CacheService[any]
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
	// Ensure a DSN s set before attempting to open the database.
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

// SchedulerService represents a service for managing scheduler.
type SchedulerService struct {
	*DB
}

// NewSchedulerService returns a new instance of SchedulerService.
func NewSchedulerService(db *DB) *SchedulerService {
	return &SchedulerService{DB: db}
}

type Schedule struct {
	Time string `json:"time"`
}

func (s *SchedulerService) Register(ctx context.Context, req scheduler.RegisterSchedulerReq) (scheduler.RegisterSchedulerRes, error) {
	i := Schedule{
		Time: req.Time,
	}
	if err := i.commit(ctx, s); err != nil {
		return scheduler.RegisterSchedulerRes{}, err
	}
	return scheduler.RegisterSchedulerRes{Status: "registered scheduler"}, nil
}

func (sc Schedule) commit(ctx context.Context, s *SchedulerService) error {
	query := `{
		q(func: type(Scheduler)) @filter(eq(name, "` + sc.Time + `")) { 
			schedulerUID as uid
		}
	}`

	res, err := upsert[Schedule](ctx, s, query, &api.Mutation{
		Cond: `@if(eq(len(schedulerUID), 0))`, //scheduler does not exist
		SetNquads: []byte(`
			_:schedule <time> "` + sc.Time + `" . .
			_:schedule <dgraph.type> "Scheduler" .
		`),
	})
	if err != nil {
		return fmt.Errorf("err while upserting scheduler: %v", err)
	}
	if len(res) != 0 {
		return errors.ConflictError(fmt.Errorf("%s exists", sc.Time))
	}
	return nil
}
