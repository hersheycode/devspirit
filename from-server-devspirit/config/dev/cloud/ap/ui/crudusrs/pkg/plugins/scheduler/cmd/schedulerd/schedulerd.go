package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"apppathway.com/examples/prodapi/pkg/plugins/scheduler"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/dgraph"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/net/grpc"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/ristretto"
	"apppathway.com/pkg/errors"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Instantiate a new type to represent our application.
	// This type lets us shared setup code with our end-to-end tests.
	m, err := NewMain()
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = m.Run()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func (m *Main) Run() (err error) {
	// Setup signal handlers.

	//open the database. This will instantiate the dgraph connection
	fmt.Println("Addr for DB", m.DB.DSN)
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db conn: %w", err)
	}

	// Initialize ristretto-backed caching for quick code generation (heavy-payloads).
	// We are using an ??? in-memory implementation but this could be changed to
	// a more robust service if we expanded out to multiple nodes.
	cacheService := ristretto.NewCacheService[any]()

	// Attach our cache service and k8s cli service to the dgraph database.
	m.DB.Cache = cacheService

	// Instantiate dgraph-backed services.
	schedulerService := dgraph.NewSchedulerService(m.DB)

	// Attach services to Main for testing.
	m.SchedulerService = schedulerService
	m.CacheService = cacheService

	// Copy configuration settings to the GRPC server.
	m.GRPCSchedulerServer.Addr = m.GRPC.Addr

	// Attach underlying services to the GRPC server.
	m.GRPCSchedulerServer.SchedulerService = schedulerService

	// Start the gRPC server.
	if err := grpc.Open(m.GRPCSchedulerServer); err != nil {
		return err
	}
	m.Close()
	return nil
}

// Close gracefully stops the program.
func (m *Main) Close() {
	if m.GRPCSchedulerServer != nil {
		m.GRPCSchedulerServer.Close()
	}
	if m.DB != nil {
		m.DB.Close()
	}
}

// Main represents the program.
type Main struct {
	Config
	// dgraph database used by dgraph service implementations.
	DB *dgraph.DB

	// GRPC server for handling GRPC communication.
	// dgraph services are attached to it before running.
	GRPCSchedulerServer *grpc.SchedulerServer

	// Services exposed for end-to-end tests.
	SchedulerService scheduler.SchedulerService

	CacheService *ristretto.CacheService[any]
}

// NewMain returns a new instance of Main.
func NewMain() (*Main, error) {
	conf, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return &Main{
		Config:              conf,
		DB:                  dgraph.New(conf.DB.DSN, conf.DB.TLS),
		GRPCSchedulerServer: grpc.NewSchedulerServer(conf.GRPC.TLS),
	}, nil
}

type Config struct {
	DB struct {
		DSN string `toml:"dsn"`
		TLS *tls.Config
	} `toml:"db"`

	GRPC struct {
		Addr string `toml:"addr"`
		TLS  *tls.Config
	} `toml:"grpc"`
}

// DefaultConfig returns a new instance of Config with defaults set.
func DefaultConfig() (Config, error) {
	conf := Config{}
	certPemRead, err := os.ReadFile(os.Getenv("CERT_FILE"))
	if err != nil {
		return conf, errors.UnexpectedError(fmt.Errorf("error loading cert file: %v", err))
	}
	certPrivKeyPEMRead, err := os.ReadFile(os.Getenv("ID_CERT_FILE"))
	if err != nil {
		return conf, errors.UnexpectedError(fmt.Errorf("error loading key file: %v", err))
	}
	serverCert, err := tls.X509KeyPair(certPemRead, certPrivKeyPEMRead)
	if err != nil {
		return conf, errors.UnexpectedError(fmt.Errorf("error loading cert from files: %v", err))
	}
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile(os.Getenv("DB_CA_FILE"))
	if err != nil {
		return conf, errors.UnexpectedError(fmt.Errorf("error loading ca file: %v", err))
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append root CA cert")
	}
	conf = Config{
		DB: struct {
			DSN string `toml:"dsn"`
			TLS *tls.Config
		}{
			DSN: os.Getenv("DB_DSN"),
			TLS: &tls.Config{
				RootCAs: rootCertPool,
			},
		},
		GRPC: struct {
			Addr string `toml:"addr"`
			TLS  *tls.Config
		}{
			Addr: os.Getenv("SCHEDULERD_ADDRESS"),
			TLS: &tls.Config{
				Certificates: []tls.Certificate{serverCert},
			},
		},
	}
	return conf, nil
}
