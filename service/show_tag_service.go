package service

import (
	"myblog/model"
	"myblog/serializer"
)

// ShowTagService 获取标签服务
type ShowTagService struct {
}

// Show 显示标签
func (service *ShowTagService) Show(id string) serializer.Response {
	var tag model.Tag
	if err := model.DB.First(&tag, id).Error; err != nil {
		return serializer.ParamErr("Tag不存在", err)
	}
	return serializer.BuildTagResponse(tag)
}
