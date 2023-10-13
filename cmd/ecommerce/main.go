package main

import (
	"log"
	"os"

	"github.com/anti-duhring/go-grpc/internal/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Conn *grpc.ClientConn
)

func main() {
	Conn, err := grpc.Dial(":8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer Conn.Close()

	r := gin.Default()
	http.SetRoutes(r)

	portEnv := os.Getenv("API_PORT")
	port := portEnv

	if portEnv == "" {
		port = "8080"
	}

	r.Run(":" + port)

	// inv := invoicer.NewInvoicerClient(Conn)

	// message := invoicer.CreateRequest{
	// 	Amount: &invoicer.Amount{
	// 		Amount:   300,
	// 		Currency: "USD",
	// 	},
	// 	From: "123",
	// 	To:   "456",
	// }

	// res, err := inv.Create(context.Background(), &message)
	// if err != nil {
	// 	log.Fatalf("could not create invoice: %s", err)
	// }

	// log.Printf("Response from Server: %s", res.Pdf)
}
