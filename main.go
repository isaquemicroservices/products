package main

import (
	"log"
	"net"

	app "github.com/isaqueveras/products-microservice/application/product"
	config "github.com/isaqueveras/products-microservice/configuration"
	inter "github.com/isaqueveras/products-microservice/interfaces/product"
	gogrpc "google.golang.org/grpc"
)

func main() {
	// Listen on port
	listen, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Creating new server
	server := gogrpc.NewServer()

	// Product registration server
	app.RegisterProductsServer(server, &inter.Server{})

	// Message of success
	log.Println("Server is running in port", config.Port)

	// Initializing server
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err.Error())
	}
}
