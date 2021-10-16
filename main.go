package main

import (
	"dateplan/config"
	"dateplan/controlls"
	"dateplan/model"
	"dateplan/servers"
	"dateplan/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	cfg:=config.GetConfig()
	err:=utils.InitLog(cfg.Log.Folder,"",time.Duration(cfg.Log.MaxAge)*time.Hour*24,time.Duration(cfg.Log.RotationTime)*time.Hour*24)
	if err!=nil{
		fmt.Println("log init fail")
		return
	}
	model.InitDB()
	servers.InitServer()
	r:=gin.Default()
	controlls.RegisterRouter(r)
	err=r.Run(":3030")
	if err != nil {
		utils.Log.Errorf("start server failed,err:%s", err)
		return
	}
}
