package grpc_client

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Conn *grpc.ClientConn

func Initialize() error {
	gatewayPort := os.Getenv("GATEWAY_PORT")

	if gatewayPort == "" {
		gatewayPort = "8089"
	}

	var err error
	Conn, err = grpc.Dial("gateway:"+gatewayPort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("could not connect: %s", err)
		return err
	}

	return nil
}
