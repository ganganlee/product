package model

import "github.com/jinzhu/gorm"

//项目数据表
type Product struct {
	gorm.Model
	Name  string `gorm:"size:125;not null" form:"name" json:"name" binding:"required,gt=10"`
	Num   uint   `gorm:"not null;default:0" form:"num" json:"num" binding:"required,numeric"`
	Image string `gorm:"size:500;not null" form:"image" json:"image" binding:"required"`
	Url   string `gorm:"size:125;not null" form:"url" json:"url" binding:"required,uri"`
	Orders []Order//定义一对多数据模型
}