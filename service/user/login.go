package user

import (
	"BlogProject/logger"
	"BlogProject/model"
	"BlogProject/pkg/e"
	"BlogProject/pkg/util"
	"BlogProject/serializer"
	"github.com/jinzhu/gorm"
)

// LoginService 登录服务
type LoginService struct {
	UserName string `form:"userName" json:"userName" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

func (service *LoginService) Login() *serializer.Response {
	var user model.User
	code := e.SUCCESS
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
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
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
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
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}
