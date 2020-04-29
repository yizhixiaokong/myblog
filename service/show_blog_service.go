package service

import (
	"myblog/model"
	"myblog/serializer"
)

// ShowBlogService 博客详情服务
type ShowBlogService struct {
}

// Show 显示博客
func (service *ShowBlogService) Show(id string) serializer.Response {
	var blog model.Blog
	if err := model.DB.First(&blog, id).Error; err != nil {
		return serializer.ParamErr("博客不存在", err)
	}
	return serializer.BuildBlogResponse(blog)
}
