package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "product.zozoo.net/Server"
	. "product.zozoo.net/common"
	"product.zozoo.net/library"
	"product.zozoo.net/model"
	"strconv"
)

//获取所有商品列表
func GetAll(c *gin.Context) {
	var product []model.Product
	Db.Find(&product)
	ResponseSuccess(c,product,"ok")
}

//修改商品
func Update(c *gin.Context) {
	var product = new(model.Product)
	//模型绑定获取参数
	err := c.ShouldBind(product)
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//判断项目id是否存在
	id,err := strconv.Atoi(c.PostForm("id"))
	if err != nil || id < 0{
		ResponseError(c, "", "参数错误")
		return
	}

	//修改
	Db.Model(product).Where("id = ?",id).Update(map[string]interface{}{"name": product.Name, "num": product.Num, "image": product.Image,"url":product.Url})
	ResponseSuccess(c, gin.H{"id": id}, "ok")
}

//新增商品
func Insert(c *gin.Context) {
	var product = new(model.Product)

	//模型绑定获取参数
	err := c.ShouldBind(product)
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//保存数据
	Db.Create(product)

	//判断数据是否保存成功
	if product.ID == 0 {
		//记录日志
		Log := library.InitLogger()
		slice, _ := json.Marshal(product)
		Log.Error("保存商品失败", zap.String("struct", string(slice)))

		Response(c, 200, 500, "", "服务器繁忙")
	} else {
		ResponseSuccess(c, gin.H{"id": product.ID}, "ok")
	}
}

//删除商品
func Delete(c *gin.Context) {
	var product = new(model.Product)
	id,err := strconv.Atoi(c.PostForm("id"))
	if err != nil || id < 0{
		ResponseError(c,"","参数错误")
		return
	}
	//删除
	Db.Unscoped().Where("id = ?",id).Delete(product)
	ResponseSuccess(c,gin.H{"id":id},"ok")
}
