package controllers

import (
	"chatOne/models"
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	web.Controller
}

// @router / [get]
func (u *UserController) UserOne() {
	iduser := u.GetString(":iduser")
	u.Ctx.WriteString("Usuario: " + iduser)
}

// @router / [put]
func (u *UserController) UpdateUser() {
	iduser := u.GetString(":iduser")
	u.Ctx.WriteString("Usuario actualizado: " + iduser)
}

// @router / [delete]
func (u *UserController) DeleteUser() {
	iduser := u.GetString(":iduser")
	u.Ctx.WriteString("Usuario Eliminado: " + iduser)
}

// retorna pagina html para ver el perfil del usuario
func (u *UserController) Usuario() {
	iduser := u.GetString(":iduser")
	u.Ctx.WriteString("Pagina de usuario " + iduser)
}

func (u *UserController) CreateUser() {
	db := models.GetDatabase()

	var user models.Usuario
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error(err.Error())
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}

	var dbuser models.Usuario
	db.Where("email = ?", user.Email).First(&dbuser)
	logs.Info("Create user dbuser ", user)
	if dbuser.Email != "" {
		u.Data["json"] = "Email ya existe"
		u.ServeJSON()
		return
	}

	//user.Password, err = GeneratePassword(user.Password)
	psw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(psw)

	// insertamos el usuario
	db.Create(&user)
	u.Data["json"] = user
	u.ServeJSON()
}


func (u *UserController) AllUser() {
	u.Data["json"] = models.Usuarios
	u.ServeJSON()
}
