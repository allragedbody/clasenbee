package controllers

import (
	"allragedbody.com/clasenbee/models"
	"encoding/json"
	"fmt"
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

//数据库操作
func createModule(b []byte) error {
	fm := &models.FileModule{}
	err := json.Unmarshal(b, fm)
	if err != nil {

	}
	err = fm.InsertFileModule()
	if err != nil {
		return err
	}
	fm.ReadFileModule()
	return nil
}

func updateModule(b []byte) error {
	fm := &models.FileModule{}
	err := json.Unmarshal(b, fm)
	if err != nil {

	}
	err = fm.InsertFileModule()
	if err != nil {
		return err
	}
	fm.ReadFileModule()
	return nil
}

func deleteModule(b []byte) error {
	fm := &models.FileModule{}
	err := json.Unmarshal(b, fm)
	if err != nil {

	}
	err = fm.InsertFileModule()
	if err != nil {
		return err
	}
	fm.ReadFileModule()
	return nil
}

// //展示模块的信息
// func (c *ModuleController) GetModeInfo() {
// 	m := FileModule{ModuleName: "CDN_LEVE2CACHE_NGINX_MAIN", ModuleChName: "CDN二级缓存nginx主配置", ModulePath: "/export/server/UCCMS/mod/nginx/cdn/nginx.mod", ModuleDesc: "二级缓存nginx主配置文件"}
// 	newMap := make(map[string]string, 20)
// 	newMap["vhostname"] = "s.jd.com.com"
// 	newMap["upstream"] = "192.168.0.1:8080\n192.168.0.2:8080\n192.168.0.3:8080\n192.168.0.4:8080\n"
// 	m.ModuleParams = newMap
// 	c.Data["json"] = m
// 	c.ServeJSON()
// }
//
//
func (c *ModuleController) CreateModuleTpl() {
	// c.Data["name"] = "失败"
	// c.Data["chname"] = "失败"
	// c.Data["path"] = "失败"
	// c.Data["desc"] = "失败"
	c.TplName = "createmodel.tpl"
}

// //创建模板
// func (c *ModuleController) CreateModuleApi(b []byte) error {
// 	createModule
// }

//创建模板
func (c *ModuleController) CreateModuleResult() {
	// var fm *models.FileModule = &models.FileModule{Name: "CDN_LEVE2CACHE_NGINX_MAIN", ChName: "CDN二级缓存nginx主配置", Path: "/export/server/UCCMS/mod/nginx/cdn/nginx.mod", Desc: "二级缓存nginx主配置文件"}

	name := c.Input().Get("name")
	chname := c.Input().Get("chname")
	path := c.Input().Get("path")
	desc := c.Input().Get("desc")
	fm := &models.FileModule{Name: name, ChName: chname, Path: path, Desc: desc}
	data, err := json.Marshal(fm)
	if err != nil {

	}
	fmt.Printf("%v", data)

	err = createModule(data)
	if err != nil {
		c.Data["result"] = "失败"
	} else {
		c.Data["result"] = "成功"
	}
	c.TplName = "createmoduleresult.tpl"
}

//创建模板
func (c *ModuleController) CreateModuleApi() {
	err := createModule(c.Ctx.Input.RequestBody)
	if err != nil {
		c.Data["json"] = "{\"ObjectId\":\"" + "Failed" + "\"}"
	} else {
		c.Data["json"] = "{\"ObjectId\":\"" + "Success" + "\"}"
	}
	c.ServeJSON()
}

//修改模板
func (c *ModuleController) ModifyModule() {

}

//删除模板
func (c *ModuleController) DeleteModule() {

}

// func (c *UserController) CreateUser() {
// 	c.Data["userid"] = "allragedbody"
// 	c.Data["tag"] = "I am a geek"
// 	c.Data["hobby"] = []string{"class", "code", "node"}
// 	c.TplName = "modulecreate.html"
// }
