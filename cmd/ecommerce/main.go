package main

import (
	"fmt"
	"os"
	"time"

	"github.com/anti-duhring/go-grpc/internal/db"
	"github.com/anti-duhring/go-grpc/internal/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Conn *grpc.ClientConn
)

func main() {
	gatewayPort := os.Getenv("GATEWAY_PORT")

	if gatewayPort == "" {
		gatewayPort = "8089"
	}

	Conn, err := grpc.Dial("gateway:"+gatewayPort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("could not connect: %s", err)
	}
	defer Conn.Close()

	time.Sleep(5 * time.Second) // Database on compose
	db.Initialize()

	if err != nil {
		fmt.Printf("could not initialize db: %s", err)
		panic(err)
	}

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
	// 	fmt.Printf("could not create invoice: %s", err)
	// 	panic(err)
	// }

	// fmt.Printf("Response from Server: %s", res.Pdf)
}
