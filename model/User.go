package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"size:125,unique" form:"name" json:"name" binding:"required,gt=3,lt=15"`
	Passwd string `gorm:"size:255" form:"passwd" json:"passwd" binding:"required,gt=6,lt=25"`
	HeadImg string `gorm:"size:125" form:"headImg" json:"headImg" binding:"required"`
}
