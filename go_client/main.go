package main

import (
	"context"
	"github.com/dalmarcogd/test_grpc/go_protos"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "0.0.0.0:3000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := go_protos.NewHelloWorldClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()
	r, err := c.SayHello(ctx, &go_protos.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}