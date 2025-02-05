package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("用户中心")
}

func (c *UserController) AddUser() {
	c.TplName = "userAdd.tpl"
}

// 处理post请求
func (c *UserController) DoAddUser() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id必须是int类型")
		return
	}
	username := c.GetString("username")
	password := c.GetString("password")
	hobby := c.GetStrings("hobby")
	fmt.Printf("爱好值：%v-------------类型：%T", hobby, hobby)
	// id 是int 类型的 需要转换成string类型
	c.Ctx.WriteString("用户中心:" + strconv.Itoa(id) + " " + username + " " + password)
}

func (c *UserController) EditUser() {
	c.TplName = "userEdit.tpl"
}

type User struct {
	// form 表示form表单提交的username和这里的Username映射，json 表示返回的字段Username使用username替换展示
	Username string   `form:"username" json:"username"`
	Password string   `form:"password" json:"password"`
	Hobby    []string `form:"hobby" json:"hobby"`
}

func (c *UserController) DoEditUser() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		c.Ctx.WriteString("post提交数据失败!")
	}
	fmt.Printf("%#v", u)
	c.Ctx.WriteString("解析post请求数据成功!")
}

func (c *UserController) GetUser() {
	u := User{
		Username: "张三",
		Password: "121312",
		Hobby:    []string{"1", "2"},
	}
	c.Data["json"] = u
	c.ServeJSON()
}
