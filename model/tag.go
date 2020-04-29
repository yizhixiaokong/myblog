package model

import (
	"github.com/jinzhu/gorm"
)

// Tag 博客模型
type Tag struct {
	gorm.Model
	Name string
}
