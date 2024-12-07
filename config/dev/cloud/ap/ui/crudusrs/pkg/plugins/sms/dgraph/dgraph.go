package dgraph

import (
	"context"
	"crypto/tls"
	"fmt"

	"time"

	"apppathway.com/examples/prodapi/pkg/plugins/sms"
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

	Cache sms.CacheService[any]
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

// SMSService represents a service for managing sms.
type SMSService struct {
	*DB
}

// NewService returns a new instance of Service.
func NewSMSService(db *DB) *SMSService {
	return &SMSService{DB: db}
}

type Message struct {
	Body string `json:"body"`
}

func (s *SMSService) Send(ctx context.Context, req sms.SendReq) (sms.SendRes, error) {
	m := Message{
		Body: req.Body,
	}
	if err := m.commit(ctx, s); err != nil {
		return sms.SendRes{}, err
	}
	return sms.SendRes{Status: "registered sms message body"}, nil
}

func (m Message) commit(ctx context.Context, s *SMSService) error {
	query := `{
		q(func: type(Message)) @filter(eq(body, "` + m.Body + `")) { 
			msgUID as uid
		}
	}`

	res, err := upsert[Message](ctx, s, query, &api.Mutation{
		Cond: `@if(eq(len(msgUID), 0))`, //msg body exists
		SetNquads: []byte(`
			_:message <body> "` + m.Body + `" . .
			_:message <dgraph.type> "Message" .
		`),
	})
	if err != nil {
		return fmt.Errorf("err while upserting message: %v", err)
	}
	if len(res) != 0 {
		return errors.ConflictError(fmt.Errorf("%s exists", m.Body))
	}
	return nil
}
