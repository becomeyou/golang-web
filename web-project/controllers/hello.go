package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "web-project/models"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Data["msg"] = "hello beego frank set value to hello views"
	c.TplName = "hello.tpl"
}
