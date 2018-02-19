package main

import (
	"net"
	"fmt"
	"log"
	"go/authentication/serverImpl"
	"google.golang.org/grpc"
	pb "go/authentication/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"errors"
)

func main() {
	log.Println("Starting service")
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := serverImpl.ServerImpl{}
	opts := []grpc.ServerOption{grpc.UnaryInterceptor(UnaryInterceptor),
	grpc.StreamInterceptor(StreamInterceptor)}

	// create a gRPC server object
	grpcServer := grpc.NewServer(opts...)
	// attach the Ping service to the server
	pb.RegisterPingServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{},error) {
	if err := authorize(ctx); err != nil {
		return nil, errors.New("Authentication failed")
	}
	return handler(ctx, req)
}

func StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := authorize(stream.Context()); err != nil {
		return errors.New("Authentication failed")
	}
	return handler(srv,stream)
}

func authorize(ctx context.Context) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md["username"]) > 0 && md["username"][0] == "admin" &&
			len(md["password"]) > 0 && md["password"][0] == "admin123" {
			return nil
		}
		log.Println("Access denied. Credentials does not match")
		return errors.New("Access denied. Credentials does not match")
	}
	log.Println("Empty metadata error")
	return errors.New("Empty metadata error")
}
