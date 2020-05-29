package main

import (
	"log"
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2"
	hello "github.com/zbrechave/micro-study/srv/proto"
	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
	"strings"
)

type Say struct {
	Client hello.SayService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error  {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.greeter", "Name cannot be blank")
	}
	response, err := s.Client.Hello(ctx, &hello.Request{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200

	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil

}

func main() {
	service := micro.NewService(micro.Name("go.micro.api.greeter"))

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client: hello.NewSayService("go.micro.srv.greeter", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

