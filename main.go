package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// Preload values used by Beego on startup
	_ "github.com/emadera52/sixty/models"
	_ "github.com/emadera52/sixty/routers"
	_ "github.com/go-sql-driver/mysql"
)

// Resgister the sixtyplus database with Beego's ORM
// IMPORTANT: Edit sixty/conf/app.conf changing 'mysqluser' and 'mysqlpass'
// to the values you use to access your MySQL database
// If the DB is not on the app's server at localhost:3306, edit
// the information between '@tcp(' and ')/' as appropriate
func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(127.0.0.1:3306)/sixtyplus"
	orm.RegisterDataBase("default", "mysql", mysqlReg)
}

func main() {
	beego.SessionOn = true
	beego.Run()
}
