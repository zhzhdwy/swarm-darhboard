package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"swarm/controller"
	"swarm/public"
)

func main() {
	err := public.ParseConfig("./config/app.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	config := public.GetConfig()
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	registrRouter(engine)
	engine.Run(config.AppHost + ":" + config.AppPort)
}

func registrRouter(engine *gin.Engine) {
	new(controller.BeeController).Router(engine)
}
