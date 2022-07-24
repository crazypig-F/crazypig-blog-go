package controller

import (
	"BlogProject/logger"
	"BlogProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTagController(c *gin.Context) {
	var listTagService service.TagService
	res := listTagService.List()
	c.JSON(http.StatusOK, res)
}

func CreateTagController(c *gin.Context) {
	var createTagService service.TagService
	if err := c.ShouldBind(&createTagService.TagForm); err == nil {
		res := createTagService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		logger.Logger.Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowTagController(c *gin.Context) {

}

func DeleteTagController(c *gin.Context) {

}

func UpdateTagController(c *gin.Context) {

}

func SearchTagController(c *gin.Context) {

}
