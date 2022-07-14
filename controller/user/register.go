package user

import (
	"BlogProject/controller"
	"BlogProject/logger"
	"BlogProject/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterController(c *gin.Context) {
	var registerService user.RegisterService
	if err := c.ShouldBind(&registerService); err == nil {
		res := registerService.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, controller.ErrorResponse(err))
		logger.Logger.Info(err)
	}
}
