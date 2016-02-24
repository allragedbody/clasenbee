package controllers

import (
	"allragedbody.com/clasenbee/models"
	"github.com/astaxie/beego"
)

type FileModule struct {
	ModuleId     int64  `json:id`
	ModuleName   string `json:name`
	ModuleChName string `json:chname`
	ModulePath   string `json:path`
	ModuleDesc   string `json:description`
	ModuleParams map[string]string
}

type ModuleController struct {
	beego.Controller
}

//展示模块的信息
func (c *ModuleController) GetModeInfo() {
	m := FileModule{ModuleName: "CDN_LEVE2CACHE_NGINX_MAIN", ModuleChName: "CDN二级缓存nginx主配置", ModulePath: "/export/server/UCCMS/mod/nginx/cdn/nginx.mod", ModuleDesc: "二级缓存nginx主配置文件"}
	newMap := make(map[string]string, 20)
	newMap["vhostname"] = "s.jd.com.com"
	newMap["upstream"] = "192.168.0.1:8080\n192.168.0.2:8080\n192.168.0.3:8080\n192.168.0.4:8080\n"
	m.ModuleParams = newMap
	c.Data["json"] = m
	c.ServeJSON()
}

//创建模板
func (c *ModuleController) CreateModule() {
	var fm *models.FileModule = &models.FileModule{Name: "CDN_LEVE2CACHE_NGINX_MAIN", ChName: "CDN二级缓存nginx主配置", Path: "/export/server/UCCMS/mod/nginx/cdn/nginx.mod", Desc: "二级缓存nginx主配置文件"}
	fm.InsertFileModule()
	fm.ReadFileModule()
}

//修改模板
func (c *ModuleController) ModifyModule() {

}

//删除模板
func (c *ModuleController) DeleteModule() {

}
