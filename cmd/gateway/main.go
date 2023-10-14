package main

import (
	"fmt"
	"net"
	"os"

	"github.com/anti-duhring/go-grpc/internal/invoicer"
	"google.golang.org/grpc"
)

func main() {
	portEnv := os.Getenv("API_PORT")
	port := portEnv

	if portEnv == "" {
		port = "8089"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("cannot create listener: %s", err)
		panic(err)
	}

	registrar := grpc.NewServer()
	service := &invoicer.Server{}

	fmt.Printf("starting server on port %v...", port)
	invoicer.RegisterInvoicerServer(registrar, service)

	err = registrar.Serve(listener)
	if err != nil {
		fmt.Printf("cannot start server: %s", err)
		panic(err)
	}
}
