package controllers

import (
	"chatOne/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

// @router / [get]
func (u *UserController) UserOne() {
	iduser := u.GetString(":iduser")
	logs.Info("iduser: ", iduser)
	u.Ctx.WriteString(iduser)
}

func (u *UserController) AllUser() {
	u.Data["json"] = models.Usuarios
	u.ServeJSON()
}
