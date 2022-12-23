package main

import (
	"github.com/sjxiang/op-panel/conf"
	"github.com/sjxiang/op-panel/models"
	"github.com/sjxiang/op-panel/router"
)


func Init() {
	conf.Init()
	models.NewDB()
}


func main() {
	Init()

	r := router.Setup()
	r.Run(":8080")
}