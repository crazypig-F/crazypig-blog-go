package serializer

type TypeVO struct {
	Id      uint   `json:"id" from:"id"`
	Name    string `json:"name" form:"name"`
	PicUrl  string `json:"picUrl" form:"picUrl"`
	BlogIds []int  `json:"blogIds" form:"blogIds"`
}
