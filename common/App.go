package common

import "fmt"

//错误处理
func HandlerErr(err error,msg string){
	if err != nil{
		fmt.Printf("%v:%v",msg,err)
	}
}
