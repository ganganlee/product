package controller

import (
	"github.com/gin-gonic/gin"
	. "product.zozoo.net/Server"
	. "product.zozoo.net/common"
	"product.zozoo.net/model"
)

//用户查询
func SelectUser(c *gin.Context) {

	//获取用户名称
	name := c.PostForm("name")
	if name == "" {
		ResponseError(c, "", "参数错误")
		return
	}

	//获取用户密码
	passWd := c.PostForm("passwd")
	if passWd == "" {
		ResponseError(c, "", "参数错误")
		return
	}

	//根据用户名查找用户是否存在
	var user = new(model.User)
	Db.Where("name=?", name).First(user)
	if user.ID == 0 {
		ResponseError(c, "", "用户不存在")
		return
	}

	//比对用户密码
}

//用户创建
func CreateUser(c *gin.Context) {

}

//比对用户密码是否相等
func ValidatePassWd(src string, passWd string) bool {

}
