package main

import (
	"net"
	"fmt"
	"log"
	"go/authentication/serverImpl"
	"google.golang.org/grpc"
	pb "go/authentication/proto"
	"google.golang.org/grpc/credentials"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"errors"
)

//grpc authentication via Transport credentials -- using CA certificates
func main() {
	log.Println("Starting service")
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1111))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := serverImpl.ServerImpl{}
	creds, err := credentials.NewServerTLSFromFile("certs1/server.crt", "certs1/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	//opts = append(opts,grpc.UnaryInterceptor(AuthInterceptor))

	// create a gRPC server object
	grpcServer := grpc.NewServer(opts...)
	// attach the Ping service to the server
	pb.RegisterPingServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("Missing metadata content")
	}
	if len(meta["authorization"]) != 1 {
		return nil, errors.New("invalid token")
	}
	if meta["authorization"][0] != "valid-token" {
		return nil, errors.New("invalid token")
	}

	return handler(ctx, req)
}
