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
}
