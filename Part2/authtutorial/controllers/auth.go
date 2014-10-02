package controllers

import (
	"authtutorial/models"
	"authtutorial/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) RegisterView() {
	this.TplNames = "register.html"
}

func (this *LoginController) Register() {
	username := this.GetString("username")
	password := this.GetString("password")
	passwordre := this.GetString("passwordre")
	test := models.RegisterForm{Username: username, Password: password, PasswordRe: passwordre}

	valid := validation.Validation{}
	b, err := valid.Valid(&test)
	if err != nil {
	}
	if !b {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
		}
	} else {
		salt := utils.GetRandomString(10)
		encodedPwd := salt + "$" + utils.EncodePassword(password, salt)

		o := orm.NewOrm()
		o.Using("default")

		user := new(models.User)
		user.Username = username
		user.Password = encodedPwd
		user.Rands = salt

		o.Insert(user)

		this.Redirect("/", 302)

	}
	this.TplNames = "register.html"
}
