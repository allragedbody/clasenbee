package main

import (
	"fmt"
	// "allragedbody.com/clasenbee/models"
	// _ "allragedbody.com/clasenbee/routers"

	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/config"
)

type ConfigWorker interface {
	GetAuthorization(c string) *DBConfig
	CreateConfig(dbc *DBConfig) *ServeConfig
	ModifyConfig() string
	DelConfig() string
}

type GitWorker interface {
	GitConfig() string
	GitAdd() string
	GitCommit() string
	GitPush() string
	GitPull() string
}

type ServerWorker interface {
	Reload(s string) string
}

type Configurator struct {
}

type Giter struct {
	giter string
}

type Processor struct {
	process string
}

type DBConfig struct {
	Role string
}

type ServeConfig struct {
	Role     string
	IP       string
	HostName string
}

var ServerConfig ServeConfig

func (w *Configurator) GetAuthorization(c string) *DBConfig {
	dbc := &DBConfig{}
	if c == "Master" {
		dbc.Role = "Master"
		return dbc
	}
	if c == "Server" {
		dbc.Role = "Server"
		return dbc
	}
	if c == "Client" {
		dbc.Role = "Client"
		return dbc
	}
	if c == "Controler" {
		dbc.Role = "Controler"
		return dbc
	}
	dbc.Role = "Other"
	return dbc
}

func (w *Configurator) CreateConfig(dbc *DBConfig) *ServeConfig {
	s := &ServeConfig{}

	s.Role = dbc.Role
	s.HostName = dbc.Role + "Server"
	s.IP = "192.168.56.103"
	return s
}
func (w *Configurator) ModifyConfig() string {
	return "Modify a Config"
}
func (w *Configurator) DelConfig() string {
	return "Del  a Config"
}

func (w *Giter) GitConfig() string {
	return "GitConfig"
}
func (w *Giter) GitAdd() string {
	return "GitAdd"
}
func (w *Giter) GitCommit() string {
	return "GitCommit"
}
func (w *Giter) GitPush() string {
	return "GitPush"
}
func (w *Giter) GitPull() string {
	return "GitPull"
}

func (w *Processor) Reload(s string) string {
	return "Reload"
}

type Server struct {
	// WebApi       WebApi
	ConfigWorker ConfigWorker
	GitWorker    GitWorker
	ServerWorker ServerWorker
}

func initConfig(sconfig *ServeConfig) {
	if sconfig.HostName != "" || sconfig.IP != "" || sconfig.Role != "" {
		ServerConfig.HostName = sconfig.HostName
		ServerConfig.IP = sconfig.IP
		ServerConfig.Role = sconfig.Role
	} else {
		// reloadConfig()
		panic("NOT Have Config")
		// fmt.Println("Has a config ,must check every 5 mimite")
	}

}

func StartApiServer(sc *ServeConfig, server *Server) {
	fmt.Printf("Start Api %v %v ", sc, server.GitWorker.GitCommit())

}

func main() {
	// conf, err := config.NewConfig("json", "conf/config.json")
	// if err != nil {
	// 	Printf("Config file load error, %v\n", err)
	// }
	// Println("Config file load Success")
	// models.MysqlLoad(conf)

	server := &Server{}
	cw := &Configurator{}
	gw := &Giter{}
	pw := &Processor{}
	server.ConfigWorker = cw
	server.GitWorker = gw
	server.ServerWorker = pw
	// a.ServerWorker = b

	dbc := server.ConfigWorker.GetAuthorization("Master")
	fmt.Println("Config Get from db :", dbc)

	sconfig := server.ConfigWorker.CreateConfig(dbc)

	initConfig(sconfig)

	fmt.Println("Start Api Worker, Config is ", ServerConfig)
	StartApiServer(&ServerConfig, server)

	// reloadConfig()
	// beego.Run()
}
