package service

import (
	"BlogProject/logger"
	"BlogProject/model"
	"BlogProject/pkg/e"
	"BlogProject/serializer"
)

type TagForm struct {
	Name string `json:"name" form:"name"`
}

type TagService struct {
	TagForm TagForm
}

func (service *TagService) List() *serializer.Response {
	var tags []model.Tag
	code := e.SUCCESS
	if err := model.DB.Find(&tags).Error; err != nil {
		code = e.ErrorDatabase
		logger.Logger.Info(err)
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	tagVOs := make([]serializer.TagVO, len(tags))
	for index, tag := range tags {
		tagVOs[index].Id = tag.ID
		tagVOs[index].Name = tag.Name
	}
	return &serializer.Response{
		Status: code,
		Data: tagVOs,
		Msg: e.GetMsg(code),
	}
}

func (service *TagService) Create() *serializer.Response {
	var tagModel model.Tag
	code := e.SUCCESS
	var count int64
	model.DB.Model(&model.User{}).Where("name=?", service.TagForm.Name).First(&tagModel).Count(&count)
	if count == 1 {
		code = e.ErrorExistTag
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建新的标签
	tagModel.Name = service.TagForm.Name
	if err := model.DB.Create(&tagModel).Error; err != nil {
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
