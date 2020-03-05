package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteBlogService 删除博客服务
type DeleteBlogService struct {
}

// Delete 删除博客
func (service *DeleteBlogService) Delete(id string) serializer.Response {
	var blog model.Blog
	if err := model.DB.First(&blog, id).Error; err != nil {
		return serializer.ParamErr("博客不存在", err)
	}
	//gorm软删除
	if err := model.DB.Delete(&blog).Error; err != nil {
		return serializer.ParamErr("博客删除失败", err)
	}
	return serializer.Response{}
}
