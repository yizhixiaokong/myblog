package service

import (
	"myblog/model"
	"myblog/serializer"
)

// CreateTagService 创建标签服务
type CreateTagService struct {
	Name string `form:"name" json:"name" `
}

// Create 创建标签
func (service *CreateTagService) Create() serializer.Response {
	tag := model.Tag{
		Name: service.Name,
	}

	if err := model.DB.Create(&tag).Error; err != nil {
		return serializer.ParamErr("Tag创建失败", err)
	}
	return serializer.BuildTagResponse(tag)
}
