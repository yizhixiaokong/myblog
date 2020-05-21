package model

import (
	"github.com/jinzhu/gorm"
)

// Blog 博客模型
type Blog struct {
	gorm.Model
	Author  string
	Title   string
	Details string
}
