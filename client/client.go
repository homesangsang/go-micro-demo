package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "go-micro-demo/gretter"
)

func main() {
	// create a new service
	service := micro.NewService(micro.Name("greeter.client"))

	// Initialise the client and parse command line flags
	service.Init()

	// Create the greeter
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rep, err := greeter.Hello(context.TODO(), &proto.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rep.Greeting)
}
