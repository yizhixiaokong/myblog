package service

import (
	"singo/model"
	"singo/serializer"
)

// UpdateBlogService 更新博客服务
type UpdateBlogService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=30"`       //标题，最少2字最大30字
	Details string `form:"details" json:"details" binding:"required,min=0,max=1000"` //内容，最少0字最大1000字
}

// Update 更新博客
func (service *UpdateBlogService) Update(id string) serializer.Response {
	var blog model.Blog
	if err := model.DB.First(&blog, id).Error; err != nil {
		return serializer.ParamErr("博客不存在", err)
	}

	blog.Title = service.Title
	blog.Details = service.Details

	if err := model.DB.Save(&blog).Error; err != nil {
		return serializer.ParamErr("博客修改失败", err)
	}
	return serializer.BuildBlogResponse(blog)
}
