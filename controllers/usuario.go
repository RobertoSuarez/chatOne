package controllers

import (
	"chatOne/models"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

// example /user
// @router	/user	[post]
func (u *UserController) Registrar() {
	var user models.Usuario
	if u.ParseForm(&user) != nil {
		u.Redirect("/v1/user", 301)
		return
	}
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Save()
	u.Redirect("/v1/user", 301)
}

// example /v1/user
// @router /user	[get]
func (u *UserController) GetPage() {
	u.TplName = "user.html"
}

// @router /login	[post]
func (u *UserController) Login() {
	if id := u.GetSession("id"); id != nil {
		u.Ctx.WriteString(fmt.Sprint("id: ", id))
		return
	}
	user := u.GetString("username")
	pass := u.GetString("password")
	ok := false
	var datauser models.Usuario
	for _, usuario := range models.Usuarios {
		if usuario.Email == user && usuario.Contraseña == pass {
			ok = true
			datauser = *usuario
			break
		}
	}

	if ok {
		u.SetSession("id", datauser.Id)
	} else {
		u.Ctx.WriteString("No existe el usuario")
		u.Ctx.Output.Status = 200
	}
	u.Redirect("/v1/users", 301)
	return
}

// @router /users	[get]
func (u *UserController) AllUser() {
	u.Data["json"] = models.Usuarios
	u.ServeJSON()
}
