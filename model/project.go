package model

import "github.com/jinzhu/gorm"

//Project 项目模型
type Project struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);"`
	Content string `gorm:"type:text;"`
	Techs   string `gorm:"type:varchar(255);"`
	PicUrl  string `gorm:"type:varchar(255);"`
	Url     string `gorm:"type:varchar(255);"`
	Type    int
}
