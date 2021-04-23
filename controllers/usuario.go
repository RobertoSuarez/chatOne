package controllers

import (
	"chatOne/models"
	"encoding/json"

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

// @router / [post]
func (u *UserController) Register() {
	iduser := u.GetString(":iduser")

	u.Ctx.WriteString(string(u.Ctx.Input.RequestBody) + iduser)
}

func (u *UserController) AllUser() {
	u.Data["json"] = models.Usuarios
	u.ServeJSON()
}

// retorna pagina html para ver el perfil del usuario
func (u *UserController) Usuario() {
	iduser := u.GetString(":iduser")
	u.Ctx.WriteString("Pagina de usuario " + iduser)
}

func (u *UserController) CreateUser() {
	var user models.Usuario
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error(err.Error())
		u.Data["json"] = err.Error()
		u.ServeJSON()
	}

	user.Save()

	u.Data["json"] = "Usuario creado correctamente"
	u.ServeJSON()

}
