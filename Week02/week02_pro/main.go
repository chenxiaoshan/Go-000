package main

import (
	"fmt"
	"week02/biz"
	_ "week02/routers"
)

// 类似这样处理  还不会用beego 就先这么处理吧
func main() {
	userDetail, err := biz.SelectUserDetailById(1)
	if err != nil {
		fmt.Printf("main: %+v\n", err)
	}
	fmt.Println(userDetail)
}