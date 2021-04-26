package controllers

import (
	"chatOne/models"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	web.Controller
}

// @router /:iduser [get]
func (u *UserController) UserOne() {
	iduser := u.GetString(":iduser")
	var user models.Usuario
	db := models.GetDatabase()

	switch u.Ctx.Input.GetData("Role") {
	case "admin":

		result := getUser(db, iduser, &user)
		if result.Error != nil {
			u.Data["json"] = models.SetError(true, result.Error.Error())
			u.ServeJSON()
		}
		u.Data["json"] = user
		u.ServeJSON()
	case "user":
		// si el usuario esta pidiendo su mismo id si se permite
		if strings.EqualFold(fmt.Sprint(u.Ctx.Input.GetData("iduser")), iduser) {
			result := getUser(db, iduser, &user)
			if result.Error != nil {
				u.Data["json"] = models.SetError(true, result.Error.Error())
				u.ServeJSON()
			}
			u.Data["json"] = user
			u.ServeJSON()
			return
		}
		u.Data["json"] = models.SetError(true, "user no permitido")
		u.ServeJSON()
	}

}

func getUser(db *gorm.DB, iduser string, user *models.Usuario) *gorm.DB {
	return db.Where("id = ?", iduser).First(&user)
}

// @router /:iduser [put]
func (u *UserController) UpdateUser() {
	var user models.Usuario
	iduser, err := strconv.Atoi(u.GetString(":iduser"))
	if err != nil {
		u.Data["json"] = models.SetError(true, err.Error())
		u.ServeJSON()
		return
	}
	user.ID = uint(iduser)
	db := models.GetDatabase()
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err != nil {
		u.Data["json"] = models.SetError(true, err.Error())
		u.ServeJSON()
		return
	}

	result := db.Model(&user).Omit("password").Updates(user)
	if result.Error != nil {
		u.Data["json"] = models.SetError(true, result.Error.Error())
		u.ServeJSON()
		return
	}

	u.Data["json"] = models.SetError(false, "Usuario actualizado")
	u.ServeJSON()

}

// @router /:iduser [delete]
func (u *UserController) DeleteUser() {
	iduser := u.GetString(":iduser")
	db := models.GetDatabase()
	result := db.Delete(&models.Usuario{}, iduser)
	if result.Error != nil {
		u.Data["json"] = models.SetError(true, result.Error.Error())
		u.ServeJSON()
		return
	}

	u.Data["json"] = models.SetError(false, iduser + " Usuario eliminado")
	u.ServeJSON()
}

// @router / [post]
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

// @router / [get]
func (u *UserController) AllUser() {
	if u.Ctx.Input.GetData("Role") == "user" {
		u.Data["json"] = models.SetError(true, "no estas permitido")
		u.ServeJSON()
		return
	}
	db := models.GetDatabase()
	var users []models.Usuario
	result := db.Find(&users)

	if result.Error != nil {
		logs.Error(result.Error)
		u.Data["json"] = models.SetError(true, result.Error.Error())
		u.ServeJSON()
		return
	}

	u.Data["json"] = users
	u.ServeJSON()
}
