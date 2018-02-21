package main

import (
	"log"
	"google.golang.org/grpc"
	pb "go/authentication/proto"
	"context"
	"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/metadata"
)

//grpc authentication via Transport credentials -- using CA certificates
func main() {
	var conn *grpc.ClientConn
	creds, err := credentials.NewClientTLSFromFile("certs1/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, errr := grpc.Dial("localhost:1111", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", errr)
	}
	defer conn.Close()
	c := pb.NewPingClient(conn)
	//md := metadata.Pairs("authorization", "valid-token")
	//ctx := metadata.NewOutgoingContext(context.Background(), md)

	response, err := c.SayHello(context.Background(), &pb.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}