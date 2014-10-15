package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int       `orm:"column(id);auto"`
	Username   string    `orm:"column(username);size(50);unique"`
	Email      string    `orm:"column(email);size(255);unique"`
	Password   string    `orm:"column(password);size(128)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
	Admin      bool      `orm:"column(admin)"`
	Rands      string    `orm:"size(10)"`
}

func init() {
	orm.RegisterModel(new(User))
}
