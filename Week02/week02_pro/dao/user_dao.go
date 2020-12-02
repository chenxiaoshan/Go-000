package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	xerrors "github.com/pkg/errors"
	"week02/model"
)

func SelectUserById(id int) (*model.User, error) {
	o := orm.NewOrm()
	user := &model.User{Id: id}
	err := o.Read(user)

	if err != nil {
		fmt.Println("No result found.")
		return nil, xerrors.Wrap(err, "no user sql:xxxx")
	}

	fmt.Println(user)
	return user, err
}
