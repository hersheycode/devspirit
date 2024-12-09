package main

import (
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/dgraph"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/net/grpc"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/ristretto"
	"apppathway.com/pkg/errors"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

type Main struct {
	Config
	DB                  *dgraph.DB
	GRPCSchedulerServer *grpc.SchedulerServer
	SchedulerService    scheduler.SchedulerService
	CacheService        *ristretto.CacheService[any]
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

func (m *Main) Close() {
	if m.GRPCSchedulerServer != nil {
		m.GRPCSchedulerServer.Close()
	}
	if m.DB != nil {
		m.DB.Close()
	}
}
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	m, err := NewMain()
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = m.Run()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
func NewMain() (*Main, error) {
	conf, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return &Main{Config: conf, DB: dgraph.New(conf.DB.DSN, conf.DB.TLS), GRPCSchedulerServer: grpc.NewSchedulerServer(conf.GRPC.TLS)}, nil
}
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
	conf = Config{DB: struct {
		DSN string `toml:"dsn"`
		TLS *tls.Config
	}{DSN: os.Getenv("DB_DSN"), TLS: &tls.Config{RootCAs: rootCertPool}}, GRPC: struct {
		Addr string `toml:"addr"`
		TLS  *tls.Config
	}{Addr: os.Getenv("SCHEDULERD_ADDRESS"), TLS: &tls.Config{Certificates: []tls.Certificate{serverCert}}}}
	return conf, nil
}
func (m *Main) Run() (err error) {
	fmt.Println("Addr for DB", m.DB.DSN)
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db conn: %w", err)
	}
	cacheService := ristretto.NewCacheService[any]()
	m.DB.Cache = cacheService
	schedulerService := dgraph.NewSchedulerService(m.DB)
	m.SchedulerService = schedulerService
	m.CacheService = cacheService
	m.GRPCSchedulerServer.Addr = m.GRPC.Addr
	m.GRPCSchedulerServer.SchedulerService = schedulerService
	if err := grpc.Open(m.GRPCSchedulerServer); err != nil {
		return err
	}
	m.Close()
	return nil
}
