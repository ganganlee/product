package model

import "github.com/jinzhu/gorm"

//订单模型
type Order struct {
	gorm.Model
	UserId    uint  `gorm:"index" form:"user_id" json:"user_id"  binding:"required,numeric"`                      //用户id
	ProductId uint  `gorm:"index" form:"product_id" json:"product_id" binding:"required,numeric"`                 //商品ID
	Status    uint8 `gorm:"size:1,default:0" form:"status" json:"status" binding:"required,numeric"` //订单状态
}

//定义订单状态常量
const (
	OrderWaite = iota
	OrderSuccess
	OrderFailed
)
