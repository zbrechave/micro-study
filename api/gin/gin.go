package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"github.com/zbrechave/micro-study/api/gin/handler"
	"log"
)



func main() {
	service := web.NewService(web.Name("go.micro.api.greeter"))

	service.Init()



	say := new(handler.Say)

	router := gin.Default()
	router.GET("/greeter", say.Anything)
	router.GET("/greeter/:name", say.Hello)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
