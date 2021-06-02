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
	_, alive, dead := bs.GetAlive()
	summary := make(map[string]int)
	summary["length"] = len(config.Nodes)
	summary["alive"] = alive
	summary["dead"] = dead
	//获取peers数量
	//peers := bs.GetPeers()
	//获取以太网地址
	//addresses := bs.GetEAddress()
	//version := bs.GetVersion()
	//port := bs.GetPort()

	ctx.HTML(http.StatusOK, "index.tmp", gin.H{
		"summary": summary,
	})
}
