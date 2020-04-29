package service

import (
	"myblog/model"
	"myblog/serializer"
)

// ListTagService 标签列表服务
type ListTagService struct {
}

// List 标签列表
func (service *ListTagService) List() serializer.Response {
	var tags []model.Tag

	if err := model.DB.Find(&tags).Error; err != nil {
		return serializer.ParamErr("数据库连接错误", err)
	}
	return serializer.BuildTagsResponse(tags)
}
