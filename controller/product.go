package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK,gin.H{
		"msg":"ok",
		"controller":"get",
	})
}

//修改商品
func Update(c *gin.Context){
	//var product = new(model.Product)
	c.JSON(http.StatusOK,gin.H{
		"msg":"ok",
		"controller":"update",
	})
}

//新增商品
func Insert(c *gin.Context){
	//var product = new(model.Product)
	c.JSON(http.StatusOK,gin.H{
		"msg":"ok",
		"controller":"insert",
	})
}

//删除商品
func Delete(c *gin.Context){
	//var product = new(model.Product)
	c.JSON(http.StatusOK,gin.H{
		"msg":"ok",
		"controller":"delete",
	})
}