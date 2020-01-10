package main

import (
	"fmt"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.helloWorld"))

	service.HandleFunc("/", helloWorldHandler)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)

	}
}

func helloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, `<html><body><h1>Hello World</h1></body></html>`)
}
