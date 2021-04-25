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
	api := beego.NewNamespace("/api/v1",

		beego.NSBefore(controllers.IsAuthorized),	// filtro de token JWT

		beego.NSInclude(&controllers.LoginController{}),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{}, // CRUD users
			),

			beego.NSNamespace("/:iduser/contacts",
				beego.NSInclude(&controllers.ContactController{},	// CRUD contacts
				),
			),
		),

		// ws
		beego.NSInclude(&controllers.ChatController{}),

	)

	beego.AddNamespace(api)
}
