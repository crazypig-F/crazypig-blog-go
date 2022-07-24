package controller

import (
	"BlogProject/logger"
	"BlogProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTypeController(c *gin.Context) {
	var listTypeService service.TypeService
	res := listTypeService.List()
	c.JSON(http.StatusOK, res)
}

func CreateTypeController(c *gin.Context) {
	var createTypeService service.TypeService
	if err := c.ShouldBind(&createTypeService.TypeForm); err == nil {
		res := createTypeService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		logger.Logger.Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowTypeController(c *gin.Context) {

}

func DeleteTypeController(c *gin.Context) {

}

func UpdateTypeController(c *gin.Context) {

}

func SearchTypeController(c *gin.Context) {

}
