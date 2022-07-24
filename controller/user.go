package controller

import (
	"BlogProject/logger"
	"BlogProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginController(c *gin.Context) {
	var loginService service.UserService
	if err := c.ShouldBind(&loginService.LoginForm); err == nil {
		res := loginService.Login()
		c.JSON(http.StatusOK, res)
	} else {
		logger.Logger.Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserRegisterController(c *gin.Context) {
	var registerService service.UserService
	if err := c.ShouldBind(&registerService.RegisterForm); err == nil {
		res := registerService.Register()
		c.JSON(http.StatusOK, res)
	} else {
		logger.Logger.Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}