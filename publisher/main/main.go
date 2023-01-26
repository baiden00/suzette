package main

import (
	"context"
	"github.com/baiden00/suzette/broker/gen"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)

	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			println(err)
		}
	}(conn)

	c := gen.NewBrokerClient(conn)
	resp, err := c.HealthCheck(context.Background(), &gen.HealthCheckRequest{
		Status: "PING",
	})

	if err != nil {
		println(err.Error())
	}
	println(resp.Status)
}
