// @Title web API
// @APIVersion 1.0.0
// @Contact electrosonix12@gmail.com
package routers

import (
	"chatOne/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	// Restringido por las cookies
	ns := beego.NewNamespace("/v1",

		beego.NSBefore(controllers.FiltrarUsuario),	// filtro para los usuarios
		beego.NSInclude(&controllers.LoginController{}),

		beego.NSRouter("/users", &controllers.UserController{}, "get:AllUser"),

		beego.NSNamespace("/user/:iduser",
			beego.NSInclude(
				&controllers.UserController{},
			),

			// toda la coleccion de contactos del usuario
			beego.NSRouter("/contacts", &controllers.ContactController{}, "get:AllContact"),

			beego.NSNamespace("/contact/:idcontact",
				beego.NSInclude(&controllers.ContactController{}),
			),
		),

		// ws
		beego.NSInclude(&controllers.ChatController{}),

	)

	beego.AddNamespace(ns)
}
