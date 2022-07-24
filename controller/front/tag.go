package front

import (
	"BlogProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListFrontTagController(c *gin.Context) {
	var listTagService service.TagService
	res := listTagService.List()
	c.JSON(http.StatusOK, res)
}
