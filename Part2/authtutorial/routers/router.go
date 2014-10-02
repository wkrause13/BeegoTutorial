package routers

import (
	"authtutorial/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.LoginController{}, "get:RegisterView;post:Register")

}
