package main

import (
	"log"
	"net"

	"github.com/isaqueveras/products-microservice/configuration"
	"google.golang.org/grpc"
)

func main() {
	var (
		listen net.Listener
		err    error
	)

	// listen on port
	if listen, err = net.Listen("tcp", configuration.Port); err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Creating new server
	server := grpc.NewServer()

	// Message of success
	log.Println("Server runing in port", configuration.Port)

	// Initializing server
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err.Error())
	}
}
