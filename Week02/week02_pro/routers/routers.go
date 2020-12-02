package routers

import (
	"github.com/astaxie/beego"
	"week02/controllers"
)

// init routers
func init() {
	beego.Router("/user", &controllers.UserController{})
}
