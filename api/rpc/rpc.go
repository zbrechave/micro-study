package main

import (
	"github.com/micro/go-micro/v2"
	proto "github.com/zbrechave/micro-study/api/rpc/proto"
	"github.com/zbrechave/micro-study/api/rpc/handler"
	hello "github.com/zbrechave/micro-study/srv/proto"
	"log"
)

func main() {
	service := micro.NewService(micro.Name("go.micro.api.greeter"))

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), &handler.Greeter{
		Client: hello.NewSayService("go.micro.srv.greeter", service.Client()),
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
