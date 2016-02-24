package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {
	c.Data["userid"] = "allragedbody"
	c.Data["tag"] = "I am a geek"
	c.Data["hobby"] = []string{"class", "code", "node"}
	c.TplName = "profile.html"
}
