package main

import (
	"github.com/sjxiang/op-panel/conf"
	"github.com/sjxiang/op-panel/router"
	// "github.com/sjxiang/op-panel/models"
	"github.com/sjxiang/op-panel/service"
)

func Init() {
	conf.Init()
	// models.NewDB()
}


func main() {
	Init()

	// 定时任务
	go service.Cron.Run()


	r := router.Setup()
	r.Run(":8080")
}