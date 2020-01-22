package main

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/test_grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	grpcPort = ":3000"
)
type Server struct{}

func (s *Server) SayHello(context context.Context, in *protos.HelloRequest) (*protos.HelloResponse, error) {
	return &protos.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	protos.RegisterHelloWorldServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	log.Printf("Listen on: %s", listen.Addr())
	grpcServer.Serve(listen)
}