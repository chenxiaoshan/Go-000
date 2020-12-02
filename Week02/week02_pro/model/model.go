package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int    `orm:"column(id)" json:"id"`
	Username string `orm:"column(username)" json:"username"`
	Company  string `orm:"column(company)" json:"company"`
}

type UserInfo struct {
	Id      int    `orm:"column(id)" json:"id"`
	UserId  int    `orm:"column(user_id)" json:"userId"`
	Address string `orm:"column(address)" json:"address"`
	Phone   string `orm:"column(phone)" json:"phone"`
}

type UserAccount struct {
	Id     int `orm:"column(id)" json:"id"`
	UserId int `orm:"column(user_id)" json:"userId"`
	Coin   int `orm:"column(coin)" json:"coin"`
}

func (user *User) TableName() string {
	return "user"
}

func (userInfo *UserInfo) TableName() string {
	return "user_info"
}

func (userAccount *UserAccount) TableName() string {
	return "user_account"
}

func (userAccount UserAccount) String() string {
	return fmt.Sprintf("id:%d  userId:%d  coin:%d", userAccount.Id, userAccount.UserId, userAccount.Coin)
}

func (user User) String() string {
	return fmt.Sprintf("id:%d userName:%s company:%s", user.Id, user.Username, user.Company)
}

func (userInfo UserInfo) String() string {
	return fmt.Sprintf("id:%d  userId:%d Address:%s phone:%s", userInfo.Id, userInfo.UserId, userInfo.Address, userInfo.Phone)
}

func init() {
	orm.RegisterModel(new(User), new(UserInfo), new(UserAccount))
}
