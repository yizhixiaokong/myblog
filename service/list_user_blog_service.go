package service

import (
	"myblog/model"
	"myblog/serializer"
)

// ListUserBlogService 当前用户博客列表服务
type ListUserBlogService struct {
}

// List 博客列表
func (service *ListUserBlogService) List(user *model.User) serializer.Response {
	var blogs []model.Blog
	var userBlogs []model.UserBlog
	if err := model.DB.Where("user_id = ?", user.ID).Find(&userBlogs).Error; err != nil {
		return serializer.ParamErr("数据库连接错误", err)
	}
	for _, userBlog := range userBlogs {
		var blog model.Blog
		if err := model.DB.Where("id = ?", userBlog.BlogID).First(&blog).Error; err != nil {
			return serializer.ParamErr("数据库连接错误", err)
		}
		blogs = append(blogs, blog)
	}

	return serializer.BuildBlogsResponse(blogs)
}
