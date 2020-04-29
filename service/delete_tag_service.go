package service

import (
	"myblog/model"
	"myblog/serializer"
)

// DeleteTagService 删除标签服务
type DeleteTagService struct {
}

// Delete 删除标签
func (service *DeleteTagService) Delete(id string) serializer.Response {
	var tag model.Tag
	if err := model.DB.First(&tag, id).Error; err != nil {
		return serializer.ParamErr("标签不存在", err)
	}
	//gorm软删除
	if err := model.DB.Delete(&tag).Error; err != nil {
		return serializer.ParamErr("标签删除失败", err)
	}
	return serializer.Response{}
}
