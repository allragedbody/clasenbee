package models

import (
	// "allragedbody/clasenbee/contorllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type FileModule struct {
	Id         int64
	Name       string
	ChName     string
	Path       string
	Desc       string
	Params     string
	CreateTime string
	UpdateTime string
}

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

func (fm *FileModule) InsertFileModule() {
	var ormFileModel orm.Ormer = orm.NewOrm()
	x, err := ormFileModel.Insert(fm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Insert Success id:%v  \n", x)
	}
}

func (fm *FileModule) ReadFileModule() {
	var ormFileModel orm.Ormer = orm.NewOrm()
	err := ormFileModel.Read(fm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Read seccess info:%v \n", fm)
	}
}
