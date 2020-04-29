package model

import (
	"github.com/jinzhu/gorm"
)

// Tag 博客模型
type BlogTag struct {
	gorm.Model
	BlogID int
	TagID  int
}
