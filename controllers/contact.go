package controllers

import (
	"chatOne/models"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)


type ContactController struct {
	web.Controller
}

// @router / [get]
func (c *ContactController) AllContacts() {
	db := models.GetDatabase()
	var contactos []models.Contacto
	userid := c.GetString(":iduser")

	// // Get all matched records
	result := db.Where("user_id = ?", userid).Find(&contactos)
	if result.Error != nil {
		c.Data["json"] = models.SetError(true, result.Error.Error())
		c.ServeJSON()
		return
	}
	c.Data["json"] = contactos
	c.ServeJSON()
}

// @router / [post]
func (c *ContactController) CreateContact() {
	var (
		iduser = c.GetString(":iduser")
		db = models.GetDatabase()
		user models.Usuario
		contacto models.Contacto
	)

	// llenar el contacto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &contacto); err != nil {
		c.Data["json"] = models.SetError(true, err.Error())
		c.ServeJSON()
		return
	}

	result := db.First(&user, "ID = ?", iduser)
	if result.Error != nil {
		c.Data["json"] = models.SetError(true, result.Error.Error())
		c.ServeJSON()
		return
	}
	contacto.UserID = user.ID
	db.Create(&contacto)

	c.Data["json"] = contacto
	c.ServeJSON()
}

// @router /:idcontact [get]
func (c *ContactController) OneContact() {
	db := models.GetDatabase()
	var (
		contact models.Contacto
		iduser = c.GetString(":iduser")
		idcontact = c.GetString(":idcontact")
	)
	result := db.First(&contact, "user_id = ? AND ID = ?", iduser, idcontact)
	if result.Error != nil {
		c.Data["json"] = models.SetError(true, result.Error.Error())
		c.ServeJSON()
		return
	}

	c.Data["json"] = contact
	c.ServeJSON()
}

// @router /:idcontact [delete]
func (c *ContactController) DeleteContact() {
	defer c.ServeJSON()
	logs.Info("delete contacto")
	iduser := c.GetString(":iduser")
	idcontacto, err := strconv.Atoi(c.GetString(":idcontact"))
	if err != nil {
		c.Data["json"] = models.SetError(true, err.Error())
		return
	}

	db := models.GetDatabase()
	var contacto models.Contacto
	contacto.ID = uint(idcontacto)
	result := db.Where("user_id = ? AND id = ?", iduser, idcontacto).Delete(&contacto)
	if result.Error != nil {
		c.Data["json"] = models.SetError(true, result.Error.Error())
		return
	}
	c.Data["json"] = models.SetError(false, "El contacto ha sido eliminado")
}

// @router /:idcontact [put]
func (c *ContactController) UpdateContact() {
	defer c.ServeJSON()

	idcontacto, err := strconv.Atoi(c.GetString(":idcontact"))
	if err != nil {
		c.Data["json"] = models.SetError(true, err.Error())
		return
	}
	var (
		iduser = c.GetString(":iduser")
		contacto models.Contacto
		db = models.GetDatabase()
	)
	contacto.ID = uint(idcontacto)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &contacto); err != nil {
		c.Data["json"] = models.SetError(true, err.Error())
		return
	}

	result := db.Model(models.Contacto{}).Where("user_id = ? AND id = ?", iduser, idcontacto).Omit("id").Updates(contacto)
	if result.Error != nil {
		c.Data["json"] = models.SetError(true, result.Error.Error())
		return
	}

	c.Data["json"] = models.SetError(false, "Contacto actualizado")
}
