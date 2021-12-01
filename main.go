package main

import (
	"log"
	"net"

	pb "github.com/isaqueveras/products-microservice/application/product"
	"google.golang.org/grpc"
)

// Server init server with gRPC
type Server struct {
	pb.UnimplementedProductsServer
}

// port of server
const port = ":9090"

func main() {
	var (
		listen net.Listener
		err    error
	)

	// listen on port
	if listen, err = net.Listen("tcp", port); err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Creating new server
	server := grpc.NewServer()

	// Registering the server on grpc
	pb.RegisterProductsServer(server, &Server{})

	// Message of success
	log.Println("Server runing in port", port)

	// Initializing server
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err.Error())
	}
}
