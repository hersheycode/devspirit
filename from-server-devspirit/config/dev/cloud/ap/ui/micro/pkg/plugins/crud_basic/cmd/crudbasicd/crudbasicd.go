package main

import (
	"apppathway.com/pkg/errors"
	"codestore.localhost/crudusrs/crud_basic/api/crudbasic"
	"codestore.localhost/crudusrs/crud_basic/api/crudbasic/dgraph"
	"codestore.localhost/crudusrs/crud_basic/api/crudbasic/net/grpc"
	"codestore.localhost/crudusrs/crud_basic/api/crudbasic/ristretto"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

type Main struct {
	Config
	DB               *dgraph.DB
	GRPCIntentServer *grpc.IntentServer
	IntentService    intent.IntentService
	CacheService     *ristretto.CacheService[any]
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
	if m.GRPCIntentServer != nil {
		m.GRPCIntentServer.Close()
	}
	if m.DB != nil {
		m.DB.Close()
	}
}
func NewMain() (*Main, error) {
	conf, err := DefaultConfig(false)
	if err != nil {
		return nil, err
	}
	return &Main{Config: conf, DB: dgraph.New(conf.DB.DSN, conf.DB.TLS), GRPCIntentServer: grpc.NewIntentServer(conf.GRPC.TLS)}, nil
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
func (m *Main) Run() (err error) {
	fmt.Println("Addr for DB", m.DB.DSN)
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db conn: %w", err)
	}
	cacheService := ristretto.NewCacheService[any]()
	m.DB.Cache = cacheService
	intentService := dgraph.NewIntentService(m.DB)
	m.IntentService = intentService
	m.CacheService = cacheService
	m.GRPCIntentServer.Addr = m.GRPC.Addr
	m.GRPCIntentServer.IntentService = intentService
	if err := grpc.Open(m.GRPCIntentServer); err != nil {
		return err
	}
	m.Close()
	return nil
}
func DefaultConfig(tlsOnDB bool) (Config, error) {
	conf := Config{}
	var dbtls *tls.Config
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
	if tlsOnDB {
		rootCertPool := x509.NewCertPool()
		pem, err := os.ReadFile(os.Getenv("DB_CA_FILE"))
		if err != nil {
			return conf, errors.UnexpectedError(fmt.Errorf("error loading ca file: %v", err))
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			log.Fatal("Failed to append root CA cert")
		}
		dbtls = &tls.Config{RootCAs: rootCertPool}
	}
	conf = Config{DB: struct {
		DSN string `toml:"dsn"`
		TLS *tls.Config
	}{DSN: os.Getenv("DB_DSN"), TLS: dbtls}, GRPC: struct {
		Addr string `toml:"addr"`
		TLS  *tls.Config
	}{Addr: os.Getenv("INTENTD_ADDRESS"), TLS: &tls.Config{Certificates: []tls.Certificate{serverCert}}}}
	return conf, nil
}
