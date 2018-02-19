package serverImpl

import (
	"golang.org/x/net/context"
	pb "go/authentication/proto"
	"log"
)

type ServerImpl struct {

}

func (s *ServerImpl) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &pb.PingMessage{Greeting: "bar"}, nil
}
