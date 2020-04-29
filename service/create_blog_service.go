package service

import (
	"myblog/model"
	"myblog/serializer"
)

// CreateBlogService 创建博客服务
type CreateBlogService struct {
	Title   string `form:"title" json:"title" binding:"required,min=1,max=30"`       //标题，最少1字最大30字
	Details string `form:"details" json:"details" binding:"required,min=0,max=2000"` //内容，最少0字最大2000字
}

// Create 创建博客
func (service *CreateBlogService) Create() serializer.Response {
	blog := model.Blog{
		Title:   service.Title,
		Details: service.Details,
	}

	if err := model.DB.Create(&blog).Error; err != nil {
		return serializer.ParamErr("博客创建失败", err)
	}
	return serializer.BuildBlogResponse(blog)
}
