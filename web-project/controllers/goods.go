package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "你好beego goods"
	c.Data["num"] = 12
	c.TplName = "goods.tpl"
}

func (c *GoodsController) DoAdd() {
	c.Ctx.WriteString("增加商品")
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// 将接收到的body参数以json格式返回出去
func (c *GoodsController) DoEdit() {
	var product Product
	err := json.NewDecoder(c.Ctx.Request.Body).Decode(&product)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Ctx.WriteString("error parsing request body")
		return
	}
	c.Data["json"] = &product
	c.ServeJSON()
}

func (c *GoodsController) DoDelete() {
	c.Ctx.WriteString("删除商品")
}
