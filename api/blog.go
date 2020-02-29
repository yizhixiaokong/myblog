package api

import (
	"singo/serializer"

	"github.com/gin-gonic/gin"
)

// CreateBlog 博客创建
func CreateBlog(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "成功",
	})
}
