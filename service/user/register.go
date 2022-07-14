package user

import (
	"BlogProject/logger"
	"BlogProject/model"
	"BlogProject/pkg/e"
	"BlogProject/serializer"
)

type RegisterService struct {
	Avatar   string `form:"avatar" json:"avatar"`
	Email    string `form:"email" json:"email"`
	UserName string `form:"userName" json:"userName" binding:"required,min=3,max=15"`
	NickName string `form:"nickName" json:"nickName" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

func (service *RegisterService) Register() *serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	//表单验证
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserName = service.UserName
	user.NickName = service.NickName
	user.Email = service.Email
	user.Avatar = service.Avatar
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logger.Logger.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logger.Logger.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
