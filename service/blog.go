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
	Tags         []uint `json:"tags" from:"tags"`
}

type BLogService struct {
	CreateBlogForm CreateBlogForm
}

//List 返回所有博客信息 但不包含博客标签 因为每一篇博客标签都需要联表查询一次数据库
// 为了减少数据库IO 在返回所有博客信息时 不返回博客标签
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

//Create 创建博客
func (service *BLogService) Create() *serializer.Response {
	var blogModel model.Blog
	var allTags []model.Tags
	var count int64
	code := e.SUCCESS
	model.DB.Model(&model.Blog{}).Where("title=?", service.CreateBlogForm.Title).First(&blogModel).Count(&count)
	if count == 1 {
		code = e.ErrorExistBlog
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建新的博客
	blogModel.Title = service.CreateBlogForm.Title
	blogModel.Content = service.CreateBlogForm.Content
	blogModel.Description = service.CreateBlogForm.Description
	blogModel.TypeId = service.CreateBlogForm.Type
	blogModel.FirstPicture = service.CreateBlogForm.FirstPicture
	blogModel.Flag = service.CreateBlogForm.Flag
	logger.Logger.Info(service.CreateBlogForm)
	if err := model.DB.Create(&blogModel).Error; err != nil {
		logger.Logger.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建博客和标签对应
	for _, tag := range service.CreateBlogForm.Tags {
		tags := model.Tags{BlogId: blogModel.ID, TagId: tag}
		allTags = append(allTags, tags)
	}
	if err := model.DB.Create(&allTags).Error; err != nil {
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

// Show 返回根据博客id查询的博客信息 包含标签id和标签名称
func (service *BLogService) Show(blogId uint) *serializer.Response {
	var blogVO serializer.BlogVO
	var tagsVOs []serializer.TagsVO
	code := e.SUCCESS
	if err := model.DB.Table("blog").Select("blog.id,description,blog.created_at,first_picture,stars,flag,title,type.id as TypeId,type.name as TypeName,content").Where("blog.id=?", blogId).Joins("left join type on blog.type_id=type.id").First(&blogVO).Error; err != nil {
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if err := model.DB.Table("tags").Select("tags.blog_id as BlogId, tags.tag_id as TagId, tag.name as TagName").Where("tags.blog_id=?", blogId).Joins("left join tag on tags.tag_id=tag.id").Scan(&tagsVOs).Error; err != nil {
		code = e.ErrorDatabase
		logger.Logger.Info(err)
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	tagIds, tagNames := make([]uint, len(tagsVOs)), make([]string, len(tagsVOs))
	for index, tagVO := range tagsVOs {
		tagIds[index] = tagVO.TagId
		tagNames[index] = tagVO.TagName
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildBlogVOWithTag(blogVO, tagIds, tagNames),
		Msg:    "success",
	}
}
