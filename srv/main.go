package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/zbrechave/micro-study/srv/handler"
	hello "github.com/zbrechave/micro-study/srv/proto"
	w_log "github.com/zbrechave/micro-study/srv/wrapper/log"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	go func() {
		grpc.DialContext(context.TODO(), "127.0.0.1:9091")
		time.Sleep(time.Second)
	}()

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(w_log.LogWrapper),
	)

	service.Init()

	hello.RegisterSayHandler(service.Server(), new(handler.Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}