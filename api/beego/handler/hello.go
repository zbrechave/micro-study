package handler

import (
	"context"

	bctx "github.com/astaxie/beego/context"
	"github.com/micro/go-micro/v2/client"
	hello "github.com/zbrechave/micro-study/srv/proto"
	"log"
)

var (
	sayService hello.SayService
)

type Say struct {}

func init()  {
	sayService = hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)
}


func (s *Say) Anything(ctx *bctx.Context) {
	log.Print("Received Say.Anything API request")
	ctx.Output.JSON(map[string]string{
		"message": "Hi, this is the Greeter API",
	}, false, true)
}

func (s *Say) Hello(ctx *bctx.Context) {
	log.Print("Received Say.Hello API request")

	name := ctx.Input.Param(":name")

	response, err := sayService.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(err, false, true)
	}

	ctx.Output.JSON(response, false, true)
}