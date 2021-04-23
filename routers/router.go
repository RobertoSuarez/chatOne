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

	// Controlamos el inicio de session
	//beego.Router("/login", &controllers.LoginController{}, "get:Page;post:IniciarSession")
	//beego.Router("/logout", &controllers.LoginController{}, "*:Logout")

	userWeb := beego.NewNamespace("/user",
		beego.NSBefore(controllers.FiltrarUsuario), // filtro para los usuarios

		beego.NSNamespace("/:iduser",
			beego.NSRouter("/", &controllers.UserController{}, "get:Usuario"),
			beego.NSRouter("/contacts", &controllers.ContactController{}, "get:Contacts"),
		),
	)

	// Restringido por las cookies
	api := beego.NewNamespace("/api/v1",

		//beego.NSBefore(controllers.FiltrarUsuario),	// filtro para los usuarios
		//beego.NSInclude(&controllers.LoginController{}),
		beego.NSInclude(&controllers.LoginController{}),

		beego.NSRouter("/users", &controllers.UserController{}, "get:AllUser;post:CreateUser"),

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

	beego.AddNamespace(api, userWeb)
}
