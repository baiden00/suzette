package main

import (
	"github.com/baiden00/suzette/broker"
)

func main() {
	println("Hello World")
	serverInstance := broker.NewServer()
	go func() {
		println("Started")
	}()
	serverInstance.Start()

}
