package router

import (
	"github.com/gin-gonic/gin"
	// "github.com/sjxiang/op-panel/middleware"
	"github.com/sjxiang/op-panel/service"
)

func registerApiRoutes(router *gin.Engine) {


	// 请求验证码

	// 用户注册

	// 用户登录
	router.POST("/login", service.Login)

	// 用户详情


	v1 := router.Group("/v1")
	// v1.Use(middleware.Auth())

	v1.GET("/systemstate", service.SystemState)

	// 定时任务
	v1.GET("/tasklist", service.TaskList)
	v1.POST("/task/add", service.TaskAdd)
	v1.DELETE("/task/del", service.TaskDelete)
	v1.PUT("/task/edit", service.TaskEdit)


	// 定时任务
	v1.GET("/jobs", service.GetJobs)
	v1.POST("/jobs", service.AddJob)
	v1.DELETE("/jobs", service.DeleteJob)
}