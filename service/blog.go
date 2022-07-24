package service

import (
	"BlogProject/logger"
	"BlogProject/model"
	"BlogProject/pkg/e"
	"BlogProject/serializer"
)

type CreateBlogForm struct {
	Content      string `json:"content" from:"content"`
	Description  string `json:"description" from:"description"`
	FirstPicture string `json:"firstPicture" from:"firstPicture"`
	Flag         string `json:"flag" from:"flag"`
	Title        string `json:"title" from:"title"`
	Type         int    `json:"type" from:"type"`
	Tags         []int  `json:"tags" from:"tags"`
}

type BLogService struct {
	CreateBlogForm CreateBlogForm
}

func (service *BLogService) List() *serializer.Response {
	code := e.SUCCESS
	var blogVOs []serializer.BlogVO
	if err := model.DB.Table("blog").Select("blog.id,description,blog.created_at,first_picture,stars,flag,title,type.id as typeId,type.name as TypeName,content").Joins("left join type on blog.id=type.id").Scan(&blogVOs).Error; err != nil {
		code = e.ErrorDatabase
		logger.Logger.Info(err)
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   blogVOs,
		Msg:    e.GetMsg(code),
	}
}

func (service *BLogService) Create() *serializer.Response {
	var blog model.Blog
	var count int64
	code := e.SUCCESS
	model.DB.Model(&model.Blog{}).Where("title=?", service.CreateBlogForm.Title).First(&blog).Count(&count)
	if count == 1 {
		code = e.ErrorExistBlog
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建新的博客
	blog.Title = service.CreateBlogForm.Title
	blog.Content = service.CreateBlogForm.Content
	blog.Description = service.CreateBlogForm.Description
	blog.TypeId = service.CreateBlogForm.Type
	blog.FirstPicture = service.CreateBlogForm.FirstPicture
	blog.Flag = service.CreateBlogForm.Flag
	logger.Logger.Info(service.CreateBlogForm)
	if err := model.DB.Create(&blog).Error; err != nil {
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

func (service *BLogService) Show(blogId uint) *serializer.Response {
	var blog model.Blog
	code := e.SUCCESS
	if err := model.DB.Where("id=?", blogId).First(&blog).Error; err != nil {
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	blogVO := serializer.BlogVO{
		Id:           blog.ID,
		Description:  blog.Description,
		Content:      blog.Content,
		FirstPicture: blog.FirstPicture,
		Title:        blog.Title,
	}
	return &serializer.Response{
		Status: code,
		Data:   blogVO,
		Msg:    "success",
	}
}
