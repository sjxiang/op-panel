package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sjxiang/op-panel/middleware"
)

func registerApiRoutes(router *gin.Engine) {


	// 请求验证码

	// 用户注册

	// 用户登录

	// 用户详情


	v1 := router.Group("/v1")
	v1.Use(middleware.Auth())

	v1.POST("/system", nil)


	
}