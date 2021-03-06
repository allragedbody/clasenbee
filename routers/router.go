package routers

import (
	"allragedbody.com/clasenbee/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/api/modinfo", &controllers.ModuleController{}, "get:GetModeInfo")
	// beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)
	beego.Router("/web/createmodule", &controllers.ModuleController{}, `get:CreateModuleTpl`)
	beego.Router("/web/createmoduleresult", &controllers.ModuleController{}, "post:CreateModuleResult")

	beego.Router("/api/createmodule", &controllers.ModuleController{}, "post:CreateModuleApi")
	beego.Router("/api/modifymodule", &controllers.ModuleController{}, "post:ModifyModule")
	beego.Router("/api/deletemodule", &controllers.ModuleController{}, "post:DeleteModule")

}
