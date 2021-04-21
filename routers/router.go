package routers

import (
	"chatOne/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSInclude(
			&controllers.ChatController{},
			&controllers.UserController{}),
	)
	beego.AddNamespace(ns)
}
