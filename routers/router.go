package routers

import (
	"allragedbody.com/clasenbee/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/modinfo", &controllers.ModuleController{}, "get:GetModeInfo")
	beego.Router("/api/createmodule", &controllers.ModuleController{}, "post:CreateModule")
	beego.Router("/api/modifymodule", &controllers.ModuleController{}, "post:ModifyModule")
	beego.Router("/api/deletemodule", &controllers.ModuleController{}, "post:DeleteModule")
}
