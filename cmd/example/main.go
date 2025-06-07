package main

import (
	"io"
	"log"
	"net"

	"github.com/PavelNen/grpc/pkg/api/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := grpc.NewServer()
	service := ExampleService{}
	example.RegisterExampleServer(server, service)

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(server)

	log.Println("gRPC server is running on port 8082")
	if err := server.Serve(lis); err != io.EOF {
		log.Fatalf("failed to serve: %v", err)
	}
}

type ExampleService struct {
	example.UnimplementedExampleServer
}