package main

import (
	"net"
	"fmt"
	"log"
	"go/authentication/serverImpl"
	"google.golang.org/grpc"
	pb "go/authentication/proto"
	"google.golang.org/grpc/credentials"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

const (
	ServerCert = "certs/server.crt"
	ServerKey  = "certs/server.key"
	CA         = "certs/ca.crt"
)

//grpc authentication via Transport credentials -- using CA certificates
func main() {
	log.Println("Starting service")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5555))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(ServerCert, ServerKey)
	if err != nil {
		log.Fatalf("could not load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CA)

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append client certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	// create a server instance
	s := serverImpl.ServerImpl{}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// create a gRPC server object
	grpcServer := grpc.NewServer(opts...)
	// attach the Ping service to the server
	pb.RegisterPingServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

