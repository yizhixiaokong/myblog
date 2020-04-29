package service

import (
	"myblog/model"
	"myblog/serializer"
)

// UpdateTagService 更新标签服务
type UpdateTagService struct {
	Name string `form:"name" json:"name" `
}

// Update 更新标签
func (service *UpdateTagService) Update(id string) serializer.Response {
	var tag model.Tag
	if err := model.DB.First(&tag, id).Error; err != nil {
		return serializer.ParamErr("标签不存在", err)
	}

	tag.Name = service.Name

	if err := model.DB.Save(&tag).Error; err != nil {
		return serializer.ParamErr("标签修改失败", err)
	}
	return serializer.BuildTagResponse(tag)
}
