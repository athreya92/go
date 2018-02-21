package main

import (
	"log"
	"google.golang.org/grpc"
	pb "go/authentication/proto"
	"context"
	"google.golang.org/grpc/credentials"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

const (
	ClientCert = "certs/client.crt"
	ClientKey  = "certs/client.key"
	CA         = "certs/ca.crt"
)


//grpc authentication via Transport credentials -- using CA certificates
func main() {
	var conn *grpc.ClientConn

	certificate, err := tls.LoadX509KeyPair(ClientCert, ClientKey)
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CA)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "server", // NOTE: This is the common-name of the server that is set while generating CA signed server certificates.
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})
	conn, errr := grpc.Dial("localhost:5555", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", errr)
	}
	defer conn.Close()
	c := pb.NewPingClient(conn)

	response, err := c.SayHello(context.Background(), &pb.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}