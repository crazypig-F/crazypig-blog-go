package model

import "github.com/jinzhu/gorm"

//Comment 评论模型
type Comment struct {
	gorm.Model
	UserId          int    `gorm:"type:bigint;"`
	Avatar          string `gorm:"type:varchar(255);"`
	Email           string `gorm:"type:varchar(255);"`
	NickName        string `gorm:"type:varchar(255);"`
	BlogId          int    `gorm:"type:bigint;"`
	ParentCommentId int    `gorm:"type:bigint;"`
	AdminComment    bool
	Content         string `gorm:"type:text;"`
}
