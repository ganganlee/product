package route

import (
	"github.com/gin-gonic/gin"
	"product.zozoo.net/controller"
)

func RegisterRoute(c *gin.Engine){
	//项目管理
	product := c.Group("/product")
	{
		product.GET("/",controller.GetAll)
		product.POST("/",controller.Insert)
		product.DELETE("/",controller.Delete)
		product.PUT("/",controller.Update)
		//获取此商品下的所有订单
		product.GET("/:id",controller.GetOrderWhitProduct)
	}

	//订单管理
	order := c.Group("/order")
	{
		order.GET("/",controller.GetOrderAll)//获取所有订单
		order.GET("/:id",controller.GetOrderById)//获取单个订单
		order.POST("/",controller.InsertOrder)//创建订单
		order.PUT("/:id",controller.UpdateOrder)//修改订单
		order.DELETE("/:id",controller.DeleteOrder)//删除订单
	}
}
