package controller

import (
	"BlogProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "传入文件出错",
		})
		return
	}
	fileSize := fileHeader.Size
	url, code := service.UpLoadFileService(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": "上传成功",
		"url":     url,
	})
}
