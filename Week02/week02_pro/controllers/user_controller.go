package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"week02/biz"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) GetUserById() {
	p1 := c.Ctx.Input.Param("id")
	id, err := strconv.Atoi(p1)
	if err != nil {
		fmt.Println("conv Atoi error")
	}
	userDetail, err := biz.SelectUserDetailById(id)
	if err != nil {
		fmt.Printf("GetUserById error: %+v\n", err)
		c.Data["json"] = nil
	} else {
		c.Data["json"] = userDetail
	}
	c.ServeJSON()
}