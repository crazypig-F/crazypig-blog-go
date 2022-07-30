package serializer

import (
	"time"
)

type BlogVO struct {
	Id           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	Description  string    `json:"description"`
	FirstPicture string    `json:"firstPicture"`
	Flag         string    `json:"flag"`
	Title        string    `json:"title"`
	TypeId       int       `json:"typeId"`
	TypeName     string    `json:"typeName"`
	Stars        int       `json:"stars"`
	Content      string    `json:"content"`
}

type BlogVOWithTag struct {
	BlogVO
	TagIds   []uint   `json:"tagIds"`
	TagNames []string `json:"tagNames"`
}

// BuildBlogVOWithTag 序列化带有tag的blog信息
func BuildBlogVOWithTag(blogVO BlogVO, tagIds []uint, tagNames []string) BlogVOWithTag {
	return BlogVOWithTag{
		BlogVO: blogVO,
		TagIds: tagIds,
		TagNames: tagNames,
	}
}
