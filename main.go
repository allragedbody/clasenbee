package main

import (
	// "allragedbody.com/clasenbee/models"
	_ "allragedbody.com/clasenbee/routers"
	// . "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/config"
)

func main() {
	// conf, err := config.NewConfig("json", "conf/config.json")
	// if err != nil {
	// 	Printf("Config file load error, %v\n", err)
	// }
	// Println("Config file load Success")
	// models.MysqlLoad(conf)
	beego.Run()
}
