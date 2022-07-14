package model

import "github.com/jinzhu/gorm"

type Type struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);"`
	PicUrl   string `gorm:"type:varchar(255);"`
	Color   string `gorm:"type:varchar(255);"`
}

