package model

import "github.com/jinzhu/gorm"

//项目数据表
type Product struct {
	gorm.Model
	Name  string `gorm:"size:125;not null" form:"name" json:"name" binding:"required,gt=10"`
	Num   uint   `gorm:"not null;default:0" form:"num" json:"num" binding:"required,numeric"`
	Image string `gorm:"size:500;not null" form:"image" json:"image" binding:"required"`
	Url   string `gorm:"size:125;not null" form:"url" json:"url" binding:"required,uri"`
}

//type ProductInterface interface {
//	//新增项目
//	Insert() uint
//	//修改项目
//	Update() bool
//	//删除项目
//	Delete() bool
//	//查询项目
//	SelectById() error
//	//查询所有项目
//	SelectAll() ([]Product,error)
//}
//
//func (p *Product)Insert() uint{
//
//	return 1
//}
//
//func (p *Product)Update() bool{
//	return false
//}
//
//func (p *Product)Delete() bool{
//	return false
//}
//
//func (p *Product)SelectById() error{
//	return nil
//}
//
//func (p *Product)SelectAll() ([]Product,error){
//	return nil,nil
//}
