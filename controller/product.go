package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	. "product.zozoo.net/Server"
	. "product.zozoo.net/common"
	"product.zozoo.net/library"
	"product.zozoo.net/model"
)

//获取所有商品列表
func GetAll(c *gin.Context) {
	//获取所有数据
	//var products = []model.Product{}
	//Server.Db.Find(products)
	/**
	c.HTML(http.StatusOK,"product.html",gin.H{
		"msg":"ok",
	})*/
	c.JSON(http.StatusOK, gin.H{
		"msg":        "ok",
		"controller": "get",
	})
}

//修改商品
func Update(c *gin.Context) {
	//var product = new(model.Product)
	c.JSON(http.StatusOK, gin.H{
		"msg":        "ok",
		"controller": "update",
	})
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
	//var product = new(model.Product)
	c.JSON(http.StatusOK, gin.H{
		"msg":        "ok",
		"controller": "delete",
	})
}
