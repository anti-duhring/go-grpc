package main

import (
	"context"
	"log"

	"github.com/anti-duhring/go-grpc/internal/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	inv := invoicer.NewInvoicerClient(conn)

	message := invoicer.CreateRequest{
		Amount: &invoicer.Amount{
			Amount:   300,
			Currency: "USD",
		},
		From: "123",
		To:   "456",
	}

	res, err := inv.Create(context.Background(), &message)
	if err != nil {
		log.Fatalf("could not create invoice: %s", err)
	}

	log.Printf("Response from Server: %s", res.Pdf)
}
