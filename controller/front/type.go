package front

import (
	"BlogProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListFrontTypeController(c *gin.Context) {
	var listTypeService service.TypeService
	res := listTypeService.List()
	c.JSON(http.StatusOK, res)
}
