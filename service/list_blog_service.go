package service

import (
	"myblog/model"
	"myblog/serializer"
)

// ListBlogService 博客列表服务
type ListBlogService struct {
}

// List 博客列表
func (service *ListBlogService) List() serializer.Response {
	var blogs []model.Blog

	if err := model.DB.Order("id desc").Limit(10).Find(&blogs).Error; err != nil {
		return serializer.ParamErr("数据库连接错误", err)
	}
	return serializer.BuildBlogsResponse(blogs)
}
