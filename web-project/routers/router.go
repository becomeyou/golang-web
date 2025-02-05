package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"web-project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})

	// 默认是get请求
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAddUser")

	// 用户
	beego.Router("/user/edit", &controllers.UserController{}, "get:EditUser")
	beego.Router("/user/doEdit", &controllers.UserController{}, "post:DoEditUser")
	beego.Router("/user/getUser", &controllers.UserController{}, "get:GetUser")

	// 商品 restful-api
	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/goods", &controllers.GoodsController{}, "post:DoAdd")
	beego.Router("/goods", &controllers.GoodsController{}, "put:DoEdit")
	beego.Router("/goods", &controllers.GoodsController{}, "delete:DoDelete")

	// 操作数据库
	beego.Router("/dbUser", &controllers.DbUserController{})
	beego.Router("/dbUser", &controllers.DbUserController{}, "post:DoAdd")
	beego.Router("/dbUser", &controllers.DbUserController{}, "delete:DoDelete")
	beego.Router("/dbUser", &controllers.DbUserController{}, "put:DoUpdate")

}
