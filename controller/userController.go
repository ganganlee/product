package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	. "product.zozoo.net/Server"
	. "product.zozoo.net/common"
	"product.zozoo.net/model"
)

//用户查询
func SelectUser(c *gin.Context) {

	//获取用户名称
	name := c.Query("name")
	if name == "" {
		ResponseError(c, "", "参数错误")
		return
	}

	//获取用户密码
	passWd := c.Query("passwd")
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
	_, err := ValidatePassWd(passWd, user.Passwd)
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	ResponseSuccess(c, gin.H{"user": user}, "ok")
}

//用户创建
func CreateUser(c *gin.Context) {
	var user = new(model.User)
	err := c.ShouldBind(user)
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//生成密码
	passWd, err := GeneratePassWd(user.Passwd)
	if err != nil {
		ResponseError(c, "", "生成密码失败")
		return
	}

	//将加密的密码保存在结构体中
	user.Passwd = string(passWd)

	//创建数据
	Db.Create(user)

	if user.ID != 0 {
		ResponseSuccess(c, gin.H{"id": user.ID}, "ok")
		return
	} else {
		ResponseError(c, "", "系统繁忙")
		return
	}
}

//比对用户密码是否相等
func ValidatePassWd(src string, passWd string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(passWd), []byte(src)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

//生成密码
func GeneratePassWd(src string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
}

//根据ID获取用户信息
func SelectUserById(id uint) *model.User {
	var user = new(model.User)
	Db.First(user, id)
	return user
}
