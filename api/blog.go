package api

import (
	"myblog/service"

	"github.com/gin-gonic/gin"
)

// CreateBlog 博客创建接口
func CreateBlog(c *gin.Context) {
	var service service.CreateBlogService
	if err := c.ShouldBind(&service); err == nil {
		user := CurrentUser(c)
		res := service.Create(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowBlog 博客详情接口
func ShowBlog(c *gin.Context) {
	var service service.ShowBlogService
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListBlog 博客列表接口
func ListBlog(c *gin.Context) {
	var service service.ListBlogService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListUserBlog 当前用户博客列表接口
func ListUserBlog(c *gin.Context) {
	var service service.ListUserBlogService
	if err := c.ShouldBind(&service); err == nil {
		user := CurrentUser(c)
		res := service.List(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateBlog 博客更新接口
func UpdateBlog(c *gin.Context) {
	var service service.UpdateBlogService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteBlog 博客详情接口
func DeleteBlog(c *gin.Context) {
	var service service.DeleteBlogService
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)

}
