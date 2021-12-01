package main

import (
	"log"
	"net"

	pb "github.com/isaqueveras/products-microservice/proto"
	"google.golang.org/grpc"
)

// Server init server with gRPC
type Server struct {
	pb.UnimplementedProductsServer
}

func main() {
	var (
		listen net.Listener
		err    error
	)

	// listen on port
	if listen, err = net.Listen("tcp", ":9090"); err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Creating new server
	server := grpc.NewServer()

	// Registering the server on grpc
	pb.RegisterProductsServer(server, Server{})

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err.Error())
	}
}
