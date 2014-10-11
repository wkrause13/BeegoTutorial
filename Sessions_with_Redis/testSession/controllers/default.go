package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	m := make(map[string]interface{})
	m["username"] = "bob"
	m["SSN"] = "000-00-0000"
	m["userId"] = 1

	v := this.GetSession("username")
	if v == nil {
		this.SetSession("username", m)
		this.Data["Website"] = "No session found, so we just set it."
	} else {
		this.Data["Website"] = v.(map[string]interface{})["SSN"]
	}

	this.TplNames = "index.tpl"
}
