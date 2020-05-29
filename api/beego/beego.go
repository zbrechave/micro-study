package main

import (
	"github.com/astaxie/beego"
	"github.com/micro/go-micro/v2/web"
	"github.com/zbrechave/micro-study/api/beego/handler"
	"log"
)

func main() {
	service := web.NewService(web.Name("go.micro.api.greeter"))

	service.Init()

	say := new(handler.Say)

	beego.Get("/greeter", say.Anything)
	beego.Get("/greeter/:name", say.Hello)

	service.Handle("/", beego.BeeApp.Handlers)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}