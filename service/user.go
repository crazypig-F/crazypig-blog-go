package service

import (
	"BlogProject/logger"
	"BlogProject/model"
	"BlogProject/pkg/e"
	"BlogProject/pkg/util"
	"BlogProject/serializer"
	"github.com/jinzhu/gorm"
)

type LoginForm struct {
	UserName string `json:"userName" form:"userName" binding:"required,min=3,max=15"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=16"`
}

type RegisterForm struct {
	Avatar   string `json:"avatar" form:"avatar" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	UserName string `json:"userName" form:"userName" binding:"required,min=3,max=15"`
	NickName string `json:"nickName" form:"nickName" binding:"required,min=3,max=15"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=16"`
}

type UserService struct {
	LoginForm    LoginForm
	RegisterForm RegisterForm
}

func (service *UserService) Login() *serializer.Response {
	var userModel model.User
	code := e.SUCCESS
	if err := model.DB.Where("user_name=?", service.LoginForm.UserName).First(&userModel).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			logger.Logger.Info(err)
			code = e.ErrorNotExistUser
			return &serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logger.Logger.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if !userModel.CheckPassword(service.LoginForm.Password) {
		code = e.ErrorNotCompare
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(userModel.ID, service.LoginForm.UserName, 0)
	if err != nil {
		logger.Logger.Info(err)
		code = e.ErrorAuthToken
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUserVO(userModel), Token: token},
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Register() *serializer.Response {
	code := e.SUCCESS
	var userModel model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", service.RegisterForm.UserName).First(&userModel).Count(&count)
	//表单验证
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//加密密码
	if err := userModel.SetPassword(service.RegisterForm.Password); err != nil {
		logger.Logger.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	userModel.UserName = service.RegisterForm.UserName
	userModel.NickName = service.RegisterForm.NickName
	userModel.Avatar = service.RegisterForm.Avatar
	userModel.Email = service.RegisterForm.Email
	if err := model.DB.Create(&userModel).Error; err != nil {
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
