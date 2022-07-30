package service

import (
	"BlogProject/logger"
	"BlogProject/model"
	"BlogProject/pkg/e"
	"BlogProject/serializer"
)

type TypeForm struct {
	Name   string `json:"name" form:"name"`
	PicUrl string `json:"picUrl" form:"picUrl"`
}

type TypeService struct {
	TypeForm TypeForm
}

func (service *TypeService) List() *serializer.Response {
	var types []model.Type
	code := e.SUCCESS
	if err := model.DB.Find(&types).Error; err != nil {
		code = e.ErrorDatabase
		logger.Logger.Info(err)
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	typeVOs := make([]serializer.TypeVO, len(types))
	for index, type_ := range types {
		typeVOs[index].Id = type_.ID
		typeVOs[index].Name = type_.Name
		typeVOs[index].PicUrl = type_.PicUrl
	}
	return &serializer.Response{
		Status: code,
		Data:   typeVOs,
		Msg:    e.GetMsg(code),
	}
}

func (service *TypeService) Create() *serializer.Response {
	var typeModel model.Type
	code := e.SUCCESS
	var count int64
	model.DB.Model(&model.User{}).Where("name=?", service.TypeForm.Name).First(&typeModel).Count(&count)
	if count == 1 {
		code = e.ErrorExistType
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建新的分类
	typeModel.Name = service.TypeForm.Name
	typeModel.PicUrl = service.TypeForm.PicUrl
	if err := model.DB.Create(&typeModel).Error; err != nil {
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
