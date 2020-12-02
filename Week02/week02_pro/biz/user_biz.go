package biz

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"week02/dao"
	"week02/model"
)

type UserDetail struct {
	*model.User
	*model.UserAccount
	*model.UserInfo
}

func SelectUserDetailById(id int) (*UserDetail, error) {
	userDetail := &UserDetail{}
	// init user
	user, err := dao.SelectUserById(id)
	if err != nil {
		// 往上抛
		return nil, err
	}
	userDetail.User = user

	// init userAccount
	userAccount, err := dao.SelectUserAccountByUserId(user.Id)
	if err != nil {
		// 往上抛
		return nil, err
	}
	userDetail.UserAccount = userAccount

	// init userInfo
	userInfo, err := dao.SelectUserInfoByUserId(user.Id)
	if errors.Is(err, orm.ErrNoRows){
		userDetail.UserInfo = nil
	}
	userDetail.UserInfo = userInfo

	return userDetail, nil
}
