package main

import (
	"os"
	"time"

	"github.com/anti-duhring/go-grpc/internal/db"
	"github.com/anti-duhring/go-grpc/internal/grpc_client"
	"github.com/anti-duhring/go-grpc/internal/http"
	"github.com/gin-gonic/gin"
)

func main() {
	err := grpc_client.Initialize()

	if err != nil {
		panic(err)
	}

	defer grpc_client.Conn.Close()

	time.Sleep(5 * time.Second) // Database on compose
	err = db.Initialize()

	if err != nil {
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
}
