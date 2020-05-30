package handler


import (
	"log"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/metadata"


	hello "github.com/zbrechave/micro-study/srv/proto"

)

type Say struct {}


func (s *Say) Hello(ctx context.Context, req *hello.Request, resp *hello.Response) error {
	log.Print("Received Say.Hello request")

	md, ok := metadata.FromContext(ctx)

	log_id, ok := md.Get("log_id")

	if !ok {
		resp.Msg = "No metadata received"
		return nil
	} else {
		log.Print("log_id is: ", log_id)
	}

	log.Printf("Received metadata %v\n", md)
	resp.Msg = fmt.Sprintf("Hello %s thanks for this %v", req.Name, md)

	resp.Msg = "Hello " + req.Name
	return nil
}
