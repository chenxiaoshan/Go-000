package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	xerrors "github.com/pkg/errors"
	"week02/model"
)

func SelectUserAccountByUserId(userId int) (*model.UserAccount, error) {
	o := orm.NewOrm()
	userAccount := &model.UserAccount{UserId: userId}

	err := o.Read(userAccount)
	if err != nil {
		return nil, xerrors.Wrap(err, "no userAccount, sql:xxxxx")
	}

	fmt.Println(userAccount)
	return userAccount, nil
}
