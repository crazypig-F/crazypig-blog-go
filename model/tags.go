package model

import (
	"github.com/jinzhu/gorm"
)

//Tags 博客标签关系模型
type Tags struct {
	gorm.Model
	BlogId uint
	TagId  uint
}
