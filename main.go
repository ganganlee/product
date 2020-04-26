package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"product.zozoo.net/Server"
	"product.zozoo.net/config"
	"product.zozoo.net/library"
	"product.zozoo.net/route"
)

//声明日志类全局变量
var Log *zap.Logger

func main()  {

	//初始化日志库
	Log = library.InitLogger()

	r := gin.Default()

	//初始化配置文件
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	//链接数据库
	err = Server.InitMysql()
	if err != nil{
		fmt.Printf("链接数据库失败 err:%v",err)
		return
	}

	//注册路由
	route.RegisterRoute(r)

	//监听端口运行项目
	r.Run(":8080")
}
