package controllers

import "github.com/beego/beego/v2/server/web"
import "github.com/beego/beego/v2/core/logs"
import "strings"
import "github.com/beego/beego/v2/server/web/context"


type LoginController struct {
	web.Controller
}

// @router /login [get]
func (l *LoginController) Page() {
	l.TplName = "user.html"
}

// @router /login [post]
func (l *LoginController) IniciarSession() {
	// crea lo cookie, de IniciarSession
	// Verifica si el usuario existe
	// Redirecciona
	username := l.GetString("username")
	password := l.GetString("password")
	if username != "" && password != "" {
		l.Ctx.SetCookie("id", username+password)
	}else {
		l.Redirect("/v1/login", 301)
		return
	}
	l.Redirect("/v1/users", 301)
}

// @router /logout [get]
func (l *LoginController) Logout() {
	// elimina las cookies
	// redirecciona
	l.Ctx.SetCookie("id", "", -1)
	l.Redirect("/v1/login", 301)
}

// Filtrar usuarios
func FiltrarUsuario(ctx *context.Context) {
	logs.Info("Filtro de usuario ", ctx.Input.URL())
	if strings.Contains(ctx.Input.URL(), "/login") {
		return
	}

	id := ctx.GetCookie("id")
	if id == "" {
		ctx.Redirect(302, "/v1/login")
	}

}
