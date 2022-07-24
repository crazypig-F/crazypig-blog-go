package serializer

import "time"

type BlogVO struct {
	Id           uint      `json:"id" form:"id"`
	CreatedAt    time.Time `json:"createdAt" from:"createdAt"`
	Description  string    `json:"description" from:"description"`
	FirstPicture string    `json:"firstPicture" from:"firstPicture"`
	Flag         string    `json:"flag" from:"flag"`
	Title        string    `json:"title" from:"title"`
	TypeId       int       `json:"typeId" from:"typeId"`
	TypeName     string    `json:"typeName" from:"typeName"`
	Stars        int       `json:"stars" from:"stars"`
	Content      string    `json:"content" from:"content"`
}
