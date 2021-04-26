package controllers

import (
	"chatOne/models"
	"encoding/json"

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
	iduser := c.GetString(":iduser")
	_ = iduser
}

// @router /:idcontact [put]
func (c *ContactController) UpdateContact() {
	iduser := c.GetString(":iduser")
	_ = iduser
}
