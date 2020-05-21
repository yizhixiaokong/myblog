package service

import (
	"myblog/model"
	"myblog/serializer"
)

// UpdateUserService 更新标签服务
type UpdateUserService struct {
	Nickname    string `form:"nickname" json:"nickname" `
	Information string `form:"information" json:"information" `
}

// Update 更新用户
func (service *UpdateUserService) Update(id string) serializer.Response {
	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return serializer.ParamErr("用户不存在", err)
	}

	user.Nickname = service.Nickname
	user.Information = service.Information

	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.ParamErr("用户修改失败", err)
	}
	return serializer.BuildUserResponse(user)
}
