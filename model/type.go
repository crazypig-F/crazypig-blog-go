package model

import "github.com/jinzhu/gorm"

type Type struct {
	gorm.Model
	Name   string `gorm:"unique;type:varchar(255);"`
	PicUrl   string `gorm:"type:varchar(255);"`
}


