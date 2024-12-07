package main

import (
	"apppathway.com/examples/prodapi/pkg/plugins/sms"
	"apppathway.com/examples/prodapi/pkg/plugins/sms/dgraph"
	"apppathway.com/examples/prodapi/pkg/plugins/sms/net/grpc"
	"apppathway.com/examples/prodapi/pkg/plugins/sms/ristretto"
	"apppathway.com/pkg/errors"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

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
type Main struct {
	Config
	DB            *dgraph.DB
	GRPCSMSServer *grpc.SMSServer
	SMSService    sms.SMSService
	CacheService  *ristretto.CacheService[any]
}

func NewMain() (*Main, error) {
	conf, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return &Main{Config: conf, DB: dgraph.New(conf.DB.DSN, conf.DB.TLS), GRPCSMSServer: grpc.NewSMSServer(conf.GRPC.TLS)}, nil
}
func (m *Main) Run() (err error) {
	fmt.Println("Addr for DB", m.DB.DSN)
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db conn: %w", err)
	}
	cacheService := ristretto.NewCacheService[any]()
	m.DB.Cache = cacheService
	smsService := dgraph.NewSMSService(m.DB)
	m.SMSService = smsService
	m.CacheService = cacheService
	m.GRPCSMSServer.Addr = m.GRPC.Addr
	m.GRPCSMSServer.SMSService = smsService
	if err := grpc.Open(m.GRPCSMSServer); err != nil {
		return err
	}
	m.Close()
	return nil
}
func (m *Main) Close() {
	if m.GRPCSMSServer != nil {
		m.GRPCSMSServer.Close()
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
	}{Addr: os.Getenv("SMSD_ADDRESS"), TLS: &tls.Config{Certificates: []tls.Certificate{serverCert}}}}
	return conf, nil
}
