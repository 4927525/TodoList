package routers

import (
	"TodoList/api"
	"TodoList/conf"
	"TodoList/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouters() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", api.Register) // 注册
		v1.POST("user/login", api.Login)       // 登录
		// 使用中间件验证登录
		auth := v1.Use(middleware.JWTAuthMiddleware())
		auth.POST("task", api.CreateTask)       // 创建清单
		auth.GET("task/:id", api.ShowTask)      // 展示某条
		auth.GET("tasks", api.ListTask)         // ListTask
		auth.PUT("task/:id", api.UpdateTask)    // 修改清单
		auth.DELETE("task/:id", api.DeleteTask) // 删除清单
	}

	r.Run(conf.AppHost)
}
