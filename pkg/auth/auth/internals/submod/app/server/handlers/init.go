package handlers

import (
	dbApi "apppathway.com/pkg/db_api/internals/project/pkg/client/api"
	dbApiMetrics "apppathway.com/pkg/db_api/internals/project/pkg/client/api_metrics"
	dbTypes "apppathway.com/pkg/db_api/internals/project/pkg/types"
	pb "apppathway.com/pkg/user/auth/internals/project/api/v1"
	"apppathway.com/pkg/user/auth/internals/submod/app/server/domain"
	serverports "apppathway.com/pkg/user/auth/internals/submod/app/server/serverports"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedAuthServer
	handlerAdaptor
}
type handlerAdaptor struct{ Services serverports.Services }

func initHandlers() handlerAdaptor {
	db := dbApi.Api{ApiMetrics: dbApiMetrics.ApiMetrics{}}
	dbChans := dbTypes.StreamChans{}
	services := domain.InitServices("dgraph", db, dbChans)
	return handlerAdaptor{Services: serverports.New(services)}
}
func Listen() {
	connStr := os.Getenv("AUTH_CONNECTION_STRING")
	opts := []grpc.ServerOption{}
	secure := true
	if secure {
		certFile := os.Getenv("CERT_FILE")
		keyFile := os.Getenv("ID_CERT_FILE")
		certPemRead, _ := os.ReadFile(certFile)
		certPrivKeyPEMRead, _ := os.ReadFile(keyFile)
		serverCert, err := tls.X509KeyPair(certPemRead, certPrivKeyPEMRead)
		if err != nil {
			log.Fatalf("Error while loading cert from files: %v", err)
			return
		}
		tlsConfig := &tls.Config{Certificates: []tls.Certificate{serverCert}}
		creds := credentials.NewTLS(tlsConfig)
		if err != nil {
			log.Fatalf("Error while loading cert from file: %v", err)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}
	lis, err := net.Listen("tcp", connStr)
	if err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
	s := grpc.NewServer(opts...)
	pb.RegisterAuthServer(s, &server{handlerAdaptor: initHandlers()})
	log.Printf("Grpc server listening at %v [dont print this in prod]\n", connStr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
