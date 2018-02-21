package main

import (
	"log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "go/authentication/proto"
)
type Login struct {
	Username string
	Password string
}

//GetRequestMetadata gets the current request metadata
func (c *Login) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"username": c.Username,
		"password": c.Password,
	}, nil
}

//If true , add certificates for TLS
func (c *Login) RequireTransportSecurity() bool {
	return false
}

//RPC credential based authentication
func main() {
	var conn *grpc.ClientConn
	auth := &Login{
		Username:"admin",
		Password:"admin123",
	}
	conn, err := grpc.Dial("localhost:2222", grpc.WithInsecure(),grpc.WithPerRPCCredentials(auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &pb.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}


