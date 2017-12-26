package rpc

import (
	"golang.org/x/net/context"
	pb "grpc-example/helloworld/message"

)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *Server) SayBye(ctx context.Context, in *pb.ByeRequest) (*pb.ByeReply, error) {
	return &pb.ByeReply{Message: "Bye " + in.Name}, nil
}
