package front

import (
	"BlogProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListFrontBlogController(c *gin.Context) {
	var listBlogService service.BLogService
	res := listBlogService.List()
	c.JSON(http.StatusOK, res)
}

func ShowFrontBlogController(c *gin.Context) {
	if blogId, err := strconv.Atoi(c.Param("id")); err == nil {
		var showFrontBlogService service.BLogService
		res := showFrontBlogService.Show(uint(blogId))
		c.JSON(http.StatusOK, res)
	}
}
