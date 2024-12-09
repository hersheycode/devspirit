package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"apppathway.com/pkg/errors"
	"apppathway.com/pkg/user/behavior"
	"apppathway.com/pkg/user/behavior/dgraph"
	"apppathway.com/pkg/user/behavior/net/grpc"
	"apppathway.com/pkg/user/behavior/recomodelcli"
	"apppathway.com/pkg/user/behavior/ristretto"
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

	if err := m.RecoModelCli.Open(); err != nil {
		return fmt.Errorf("cannot open recomodelcli conn: %w", err)
	}

	//open the database. This will instantiate the dgraph connection
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db conn: %w", err)
	}

	// Initialize ristretto-backed caching for quick code generation (heavy-payloads).
	// We are using an ??? in-memory implementation but this could be changed to
	// a more robust service if we expanded out to multiple nodes.
	cacheService := ristretto.NewCacheService[any]()

	recomodelCliService := recomodelcli.NewRecoModelCliService(m.RecoModelCli)

	// Attach our cache service and k8s cli service to the dgraph database.
	m.DB.Cache = cacheService
	m.DB.RecoModelCli = recomodelCliService

	// Instantiate dgraph-backed services.
	behaviorService := dgraph.NewBehaviorService(m.DB)

	// Attach services to Main for testing.
	m.BehaviorService = behaviorService
	m.CacheService = cacheService
	m.RecoModelCliService = recomodelCliService

	// Copy configuration settings to the GRPC server.
	m.GRPCServer.Addr = m.GRPC.Addr

	// Attach underlying services to the GRPC server.
	m.GRPCServer.BehaviorService = behaviorService

	// Start the gRPC server.
	if err := m.GRPCServer.Open(); err != nil {
		return err
	}
	m.Close()
	return nil
}

// Close gracefully stops the program.
func (m *Main) Close() {
	if m.GRPCServer != nil {
		m.GRPCServer.Close()
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
	GRPCServer *grpc.Server

	// Services exposed for end-to-end tests.
	BehaviorService behavior.BehaviorService

	RecoModelCli *recomodelcli.RecoModelCli

	CacheService *ristretto.CacheService[any]

	RecoModelCliService *recomodelcli.RecoModelCliService
}

// NewMain returns a new instance of Main.
func NewMain() (*Main, error) {
	conf, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return &Main{
		Config:       conf,
		DB:           dgraph.New(conf.DB.DSN, conf.DB.TLS),
		RecoModelCli: recomodelcli.New(),
		GRPCServer:   grpc.NewServer(conf.GRPC.TLS),
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
			Addr: os.Getenv("BEHAVIORD_ADDRESS"),
			TLS: &tls.Config{
				Certificates: []tls.Certificate{serverCert},
			},
		},
	}
	return conf, nil
}
