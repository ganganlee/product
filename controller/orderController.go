package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "product.zozoo.net/Server"
	. "product.zozoo.net/common"
	"product.zozoo.net/library"
	. "product.zozoo.net/model"
	"strconv"
)

//新增订单
func InsertOrder(c *gin.Context) {
	//获取参数
	var order = new(Order)
	err := c.ShouldBind(order)

	//字段验证
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//保存数据
	Db.Create(order)
	if order.ID == 0 {
		//记录日志
		Log := library.InitLogger()
		slice, _ := json.Marshal(order)
		Log.Error("订单添加失败", zap.String("struct", string(slice)))

		ResponseError(c, gin.H{"code": 501}, "数据库繁忙")
		return
	} else {
		ResponseSuccess(c, gin.H{"id": order.ID}, "ok")
	}
}

//修改订单
func UpdateOrder(c *gin.Context) {
	//获取参数
	var order = new(Order)
	err := c.ShouldBind(order)

	//字段验证
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//获取订单Id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//修改
	Db.Model(order).Where("id = ?", id).Update(map[string]interface{}{"user_id": order.UserId, "product_id": order.ProductId, "status": order.Status})
	ResponseSuccess(c, gin.H{"id": id}, "ok")
}

//删除订单
func DeleteOrder(c *gin.Context) {
	//获取订单Id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//查找订单是否存在
	var order = new(Order)
	Db.First(order, id)
	if order.ID == 0 {
		ResponseError(c, "", "订单不存在")
		return
	}

	//删除
	Db.Unscoped().Delete(&order)
	ResponseSuccess(c, gin.H{"id": id}, "ok")
}

//获取所有订单
func GetOrderAll(c *gin.Context) {
	//实例化订单模型
	var orders []Order
	Db.Find(orders)
	ResponseSuccess(c, orders, "ok")
}

//获取单个订单
func GetOrderById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//实例化订单表
	var order = new(Order)
	Db.First(order, id)
	if order.ID == 0 {
		ResponseError(c, "", "数据不存在")
	} else {
		ResponseSuccess(c, order, "ok")
	}
}

//根据商品id获取所有订单
func GetOrderWhitProduct(c *gin.Context){
	//获取商品ID
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		ResponseError(c, "", "参数错误")
		return
	}

	//实例化商品模型
	var product Product
	product.ID = uint(id)
	var Orders []Order
	Db.Model(product).Related(Orders)

	ResponseSuccess(c,Orders,"ok")
}
