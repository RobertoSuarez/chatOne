package controllers

import (
	"encoding/json"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/google/uuid"
)


type LoginController struct {
	web.Controller
}

// TODO: /login debe de resivir un json con las creadenciales,
// y retornar un json con un token para que el js cree la cookies

// @router /login [post]
func (l *LoginController) IniciarSession() {
	var login struct {
		Email string`json:"email"`
		Password	string`json:"password"`
		Token 	string`json:"token"`
	}
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &login)
	if err != nil {
		l.Data["json"] = err.Error()
		l.ServeJSON()
		return
	}

	login.Token = uuid.New().String()
	l.Data["json"] = login
	l.ServeJSON()
}

// Filtrar usuarios
func FiltrarUsuario(ctx *context.Context) {
	logs.Info("Filtro de usuario ", ctx.Input.URL())
	if strings.Contains(ctx.Input.URL(), "/login") {
		return
	}
	// si no se tiene la cookies id, entonces
	// se redireciona a login, para iniciar session
	id := ctx.GetCookie("id")
	if id == "" {
		ctx.Redirect(302, "/login")
	}

}
