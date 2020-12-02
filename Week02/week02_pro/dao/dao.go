package dao

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql", "root:20201125@tcp(113.31.104.168:3306)/go_train_camp?charset=utf8")
	if err != nil {
		panic(errors.New("conn database error"))
		return
	}
	fmt.Println("conn database success")
}

