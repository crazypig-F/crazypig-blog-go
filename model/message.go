package model

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	NickName string `gorm:"type:varchar(255);"`
	Avatar   string `gorm:"type:varchar(255);"`
	Content  string `gorm:"type:varchar(255);"`
	UserID   int    `gorm:"type:varchar(255);"`
}
