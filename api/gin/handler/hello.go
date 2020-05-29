package handler

import (
	"context"

	"github.com/gin-gonic/gin"
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


func (s *Say) Anything(c *gin.Context)  {
	log.Print("Received Say.Anything API request")

	c.JSON(200, map[string]string{
		"message": "Hi, this is the greeter api",
	})
}

func (s *Say) Hello(c *gin.Context)  {
	log.Print("Received Say.Hello API request")

	name := c.Param("name")

	response, err := sayService.Hello(context.TODO(), &hello.Request{
		Name: name,
	})
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}
