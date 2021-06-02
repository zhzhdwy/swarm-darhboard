package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swarm/public"
	"swarm/service"
)

type BeeController struct {
}

func (bc *BeeController) Router(engine *gin.Engine) {
	// 获取验证码
	engine.GET("/", bc.nodeStatus)
}

func (uc *BeeController) nodeStatus(ctx *gin.Context) {
	config := public.GetConfig()
	bs := service.BeeService{
		Nodes: config.Nodes,
	}
	peers := bs.GetPeers()

	ctx.HTML(http.StatusOK, "index.tmp", gin.H{
		"title": peers,
	})
}
