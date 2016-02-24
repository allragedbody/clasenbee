package models

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db := beego.AppConfig.String("MysqlDB::db")
	if db == "mysql" {
		user := beego.AppConfig.String("MysqlDB::user")
		pass := beego.AppConfig.String("MysqlDB::pass")
		name := beego.AppConfig.String("MysqlDB::name")
		url := beego.AppConfig.String("MysqlDB::url")
		mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pass, url, name)
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", mysqlConnStr, 300)
	}

	orm.RegisterModel(new(FileModule))
	orm.RunSyncdb("default", false, true)
}

// func MysqlLoad(conf config.Configer) {
// 	db := conf.String("MysqlDB::db")
// 	if db == "mysql" {
// 		// user := conf.String("MysqlDB::user")
// 		pass := conf.String("MysqlDB::pass")
// 		name := conf.String("MysqlDB::name")
// 		user := conf.String("MysqlDB::user")
// 		url := conf.String("MysqlDB::url")
// 		mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pass, url, name)
// 		orm.RegisterDriver("mysql", orm.DRMySQL)
// 		orm.RegisterDataBase("default", "mysql", mysqlConnStr, 300)
// 	}

// 	orm.RegisterModel(new(FileModule))
// 	orm.RunSyncdb("default", false, true)
// }
