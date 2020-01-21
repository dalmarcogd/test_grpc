package main

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/test_grpc/go_protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	grpcPort = ":50051"
)
type Server struct{}

func (s *Server) SayHello(context context.Context, in *go_protos.HelloRequest) (*go_protos.HelloResponse, error) {
	return &go_protos.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	go_protos.RegisterHelloWorldServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}