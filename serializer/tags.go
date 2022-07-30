package serializer

type TagsVO struct {
	BlogId uint `json:"blogId" json:"blogId"`
	TagId  uint `json:"tagId" form:"tagId"`
	TagName string `json:"tagName"`
}
