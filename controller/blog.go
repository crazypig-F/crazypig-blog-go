package controller

import (
	"BlogProject/logger"
	"BlogProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListBlogController(c *gin.Context) {
	var listBlogService service.BLogService
	res := listBlogService.List()
	c.JSON(http.StatusOK, res)
}

func CreateBlogController(c *gin.Context) {
	var createBlogService service.BLogService
	if err := c.ShouldBind(&createBlogService.CreateBlogForm); err == nil {
		res := createBlogService.Create()
		c.JSON(http.StatusOK, res)
	}else{
		logger.Logger.Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowBlogController(c *gin.Context) {

}

func DeleteBlogController(c *gin.Context) {

}

func UpdateBlogController(c *gin.Context) {

}
func SearchBlogController(c *gin.Context) {

}
