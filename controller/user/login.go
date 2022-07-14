package user

import (
	"BlogProject/controller"
	"BlogProject/logger"
	"BlogProject/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginController(c *gin.Context) {
	var loginService user.LoginService
	if err := c.ShouldBind(&loginService); err == nil {
		res := loginService.Login()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, controller.ErrorResponse(err))
		logger.Logger.Info(err)
	}
}
