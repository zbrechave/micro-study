package handler

import (
	"context"
	hello "github.com/zbrechave/micro-study/srv/proto"
	proto "github.com/zbrechave/micro-study/api/rpc/proto"
	"log"
)

type Greeter struct {
	Client hello.SayService
}


func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Greeter.Hello API request")

	response, err := g.Client.Hello(ctx, &hello.Request{
		Name: req.Name,
	})

	if err != nil {
		return err
	}

	// set api response
	rsp.Msg = response.Msg
	return nil
}