package model

import (
	"github.com/jinzhu/gorm"
)

//Essay 随笔模型
type Essay struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);"`
	Content string `gorm:"type:text;"`
	Image   string `gorm:"type:varchar(255);"`
	Praise  int    `gorm:"type:bigint;"`
	Color   string `gorm:"type:varchar(255);"`
}
