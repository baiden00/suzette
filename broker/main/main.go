package main

import (
	"github.com/baiden00/suzette/broker"
	"github.com/baiden00/suzette/broker/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	println("Hello World")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	gen.RegisterBrokerServer(s, &broker.Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
