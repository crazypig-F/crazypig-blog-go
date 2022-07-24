package serializer

import "BlogProject/model"

type UserVO struct {
	UserName      string `json:"useName" form:"useName"`
	NickName      string `json:"nickName" form:"nickName"`
	Avatar        string `json:"avatar" form:"avatar"` //头像
	Administrator bool   `json:"administrator" form:"administrator"`
	CreateAt      int64  `json:"create_at" form:"create_at"` // 创建
}

//BuildUserVO 序列化用户
func BuildUserVO(user model.User) UserVO {
	return UserVO{
		UserName: user.UserName,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Administrator: user.Administrator,
		CreateAt: user.CreatedAt.Unix(),
	}
}
