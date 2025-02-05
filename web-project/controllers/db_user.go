package controllers

import (
	// 注意这里的bee-project是go.mod里面的名称 而不是项目名称
	beego "github.com/beego/beego/v2/server/web"
	"web-project/models"
)

type DbUserController struct {
	beego.Controller
}

func (c *DbUserController) DoAdd() {
	models.Add()
	c.Ctx.WriteString("保存数据成功!")
}

func (c *DbUserController) DoDelete() {
	models.Delete()
	c.Ctx.WriteString("删除数据成功!")
}

func (c *DbUserController) DoUpdate() {
	models.Update()
	c.Ctx.WriteString("更新数据成功!")
}

func (c *DbUserController) Get() {
	models.Find()
	c.Ctx.WriteString("查询数据成功!")
}
