package routers

import (
	"authtutorial/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:LoginView;post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/register", &controllers.LoginController{}, "get:RegisterView;post:Register")
	beego.Router("/secret", &controllers.LoginController{}, "get:SecretView")

	beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)

}
