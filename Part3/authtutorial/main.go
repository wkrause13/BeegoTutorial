package main

import (
	_ "authtutorial/models"
	_ "authtutorial/routers"
	_ "github.com/astaxie/beego/session/redis"

	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@(127.0.0.1:3306)/authtutorial")

}
func main() {

	beego.SessionOn = true

	name := "default"
	// Whether to drop table and re-create.
	force := false
	// Print log.
	verbose := false
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()
}
