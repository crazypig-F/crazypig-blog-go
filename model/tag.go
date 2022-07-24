package model

import (
	"github.com/jinzhu/gorm"
)

//Tag 标签模型
type Tag struct {
	gorm.Model
	Name   string `gorm:"unique;type:varchar(255);"`
}