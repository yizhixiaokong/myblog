package api

import (
	"myblog/service"

	"github.com/gin-gonic/gin"
)

// CreateTag 标签创建接口
func CreateTag(c *gin.Context) {
	var service service.CreateTagService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowTag 获取标签接口
func ShowTag(c *gin.Context) {
	var service service.ShowTagService
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListTag 标签列表接口
func ListTag(c *gin.Context) {
	var service service.ListTagService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateTag 标签更新接口
func UpdateTag(c *gin.Context) {
	var service service.UpdateTagService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteTag 标签详情接口
func DeleteTag(c *gin.Context) {
	var service service.DeleteTagService
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)

}
