package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"week02/model"
)

func SelectUserInfoByUserId(id int) (*model.UserInfo, error) {
	o := orm.NewOrm()
	userInfo := &model.UserInfo{UserId: 1}

	err := o.Read(userInfo)

	if err != nil {
		fmt.Println("No result found.")
		return nil, err
	}

	fmt.Println(userInfo.Id, userInfo.UserId, userInfo.Address, userInfo.Phone)
	return userInfo, err

}
