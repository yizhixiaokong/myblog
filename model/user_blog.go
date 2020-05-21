package model

import (
	"github.com/jinzhu/gorm"
)

// Tag 博客模型
type UserBlog struct {
	gorm.Model
	UserID int
	BlogID int
}
