package handler


import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	hello "github.com/zbrechave/micro-study/srv/proto"

)

type Say struct {}


func (s *Say) Hello(ctx context.Context, req *hello.Request, resp *hello.Response) error {
	log.Log("Received Say.Hello request")
	resp.Msg = "Hello " + req.Name
	return nil
}
