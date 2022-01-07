package api

import (
	"TodoList/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册
func Register(c *gin.Context) {
	var registerService service.LoginOrRegisterService
	// 判断传来的参数是否合法
	if err := c.ShouldBind(&registerService); err == nil {
		// 注册
		res := registerService.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 登录
func Login(c *gin.Context) {
	var loginService service.LoginOrRegisterService
	// 判断传来的参数是否合法
	if err := c.ShouldBind(&loginService); err == nil {
		// 登录
		res := loginService.Login()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
