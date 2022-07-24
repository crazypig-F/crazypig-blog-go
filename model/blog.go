package model

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Appreciate     int    `gorm:"type:bigint;"`
	Commentable    int    `gorm:"type:bigint;"`
	Content        string `gorm:"type:text;"`
	Description    string `gorm:"type:varchar(255);"`
	FirstPicture   string `gorm:"type:varchar(255);"`
	Flag           string `gorm:"type:varchar(255);"`
	Published      bool
	Recommend      bool
	ShareStatement bool
	Title          string `gorm:"unique;type:varchar(255);"`
	Stars          int
	TypeId         int `gorm:"type:bigint;"`
	UserId         int `gorm:"type:bigint;"`
}
