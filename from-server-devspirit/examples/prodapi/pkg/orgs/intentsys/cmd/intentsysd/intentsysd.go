package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"apppathway.com/examples/prodapi/pkg/orgs/intentsys"
	"apppathway.com/examples/prodapi/pkg/orgs/intentsys/net/grpc"
	"apppathway.com/examples/prodapi/pkg/orgs/intentsys/ristretto"
	"apppathway.com/examples/prodapi/pkg/plugins/intent/api/intentpb"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb"
	"apppathway.com/examples/prodapi/pkg/plugins/sms/api/smspb"

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
	// Attach services to Main for testing.
	m.IntentService = grpc.NewIntentClient(m.Intent.Addr, m.Intent.Cert)
	m.SchedulerService = grpc.NewSchedulerClient(m.Scheduler.Addr, m.Scheduler.Cert)
	m.SMSService = grpc.NewSMSClient(m.SMS.Addr, m.SMS.Cert)

	// Copy configuration settings to the GRPC server.
	m.GRPCIntentSysServer.Addr = m.GRPC.Addr

	// Attach underlying services to the GRPC server.
	m.GRPCIntentSysServer.IntentService = m.IntentService
	m.GRPCIntentSysServer.SchedulerService = m.SchedulerService
	m.GRPCIntentSysServer.SMSService = m.SMSService

	// Start the gRPC server.
	if err := grpc.Open(m.GRPCIntentSysServer); err != nil {
		return err
	}
	m.Close()
	return nil
}

// Close gracefully stops the program.
func (m *Main) Close() {
	if m.GRPCIntentSysServer != nil {
		m.GRPCIntentSysServer.Close()
	}
}

// Main represents the program.
type Main struct {
	Config
	// GRPC server for handling GRPC communication.
	// dgraph services are attached to it before running.
	GRPCIntentSysServer *grpc.IntentSysServer

	// Services exposed for end-to-end tests.
	IntentSysService intentsys.IntentSysService

	IntentService    intentpb.IntentClient
	SchedulerService schedulerpb.SchedulerClient
	SMSService       smspb.SMSClient

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
		GRPCIntentSysServer: grpc.NewIntentSysServer(conf.GRPC.TLS),
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

	SMS struct {
		Addr string `toml:"addr"`
		Cert string `toml:"cert"`
	} `toml:"sms"`

	Intent struct {
		Addr string `toml:"addr"`
		Cert string `toml:"cert"`
	} `toml:"intent"`

	Scheduler struct {
		Addr string `toml:"addr"`
		Cert string `toml:"cert"`
	} `toml:"scheduler"`
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

	conf = Config{
		GRPC: struct {
			Addr string `toml:"addr"`
			TLS  *tls.Config
		}{
			Addr: os.Getenv("INTENTSYSD_ADDRESS"),
			TLS: &tls.Config{
				Certificates: []tls.Certificate{serverCert},
			},
		},
		SMS: struct {
			Addr string `toml:"addr"`
			Cert string `toml:"cert"`
		}{
			Addr: os.Getenv("SMSD_ADDRESS"),
			Cert: os.Getenv("CA_FILE"),
		},
		Intent: struct {
			Addr string `toml:"addr"`
			Cert string `toml:"cert"`
		}{
			Addr: os.Getenv("INTENTD_ADDRESS"),
			Cert: os.Getenv("CA_FILE"),
		},
		Scheduler: struct {
			Addr string `toml:"addr"`
			Cert string `toml:"cert"`
		}{
			Addr: os.Getenv("SCHEDULERD_ADDRESS"),
			Cert: os.Getenv("CA_FILE"),
		},
	}
	return conf, nil
}
