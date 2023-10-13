package main

import (
	"fmt"
	"log"
	"net"

	"github.com/anti-duhring/go-grpc/internal/invoicer"
	"google.golang.org/grpc"
)

func main() {
	port := ":8089"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	registrar := grpc.NewServer()
	service := &invoicer.Server{}

	fmt.Printf("starting server on port %v...", port)
	invoicer.RegisterInvoicerServer(registrar, service)

	err = registrar.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
