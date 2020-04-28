package route

import (
	"github.com/gin-gonic/gin"
	"product.zozoo.net/controller"
)

func RegisterRoute(c *gin.Engine){
	product := c.Group("/product")
	{
		product.GET("/",controller.GetAll)
		product.POST("/",controller.Insert)
		product.DELETE("/",controller.Delete)
		product.PUT("/",controller.Update)
	}

	order := c.Group("/order")
	{
		order.GET("/",controller.GetOrderAll)//获取所有订单
		order.GET("/order/:id",controller.GetOrderById)//获取单个订单
		order.POST("/",controller.InsertOrder)//创建订单
		order.PUT("/",controller.UpdateOrder)//修改订单
		order.DELETE("/:id",controller.DeleteOrder)//删除订单
		order.GET("/product/:id",controller.GetOrderWhitProduct)//获取此商品下所有订单
	}
}
