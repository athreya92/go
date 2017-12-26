package main

import (
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	//-+
	// "google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "grpc-example/helloworld/service"
	"grpc-example/rpc"
)

const (
	port = ":50051"
)


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	greeterService := rpc.Server{}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterService)
//	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
