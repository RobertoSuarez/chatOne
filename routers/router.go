package routers

import (
	"chatOne/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/websocket",
		beego.NSInclude(&controllers.ChatController{}),
	)
	beego.AddNamespace(ns)
}
