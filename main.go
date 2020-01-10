package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	greeter "go-micro-demo/gretter"
	"log"
)

const serviceName string = "go.micro.greeter"

type Greeter struct {
	Id string `json: 'Id'`
}

func (g *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	log.Print("Received greeter.client API request")
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name(serviceName))

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	_ = greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
