package controllers

import (
	"github.com/beego/beego/v2/server/web"
)


type ContactController struct {
	web.Controller
}

// @router / [get]
func (c *ContactController) AllContacts() {
	c.Data["json"] = "Todos los contactos de " + c.GetString(":iduser")
	c.ServeJSON()
}

// @router / [post]
func (c *ContactController) CreateContact() {
	iduser := c.GetString(":iduser")
	_ = iduser
}

// @router /:idcontact [get]
func (c *ContactController) OneContact() {
	iduser := c.GetString(":iduser")
	_ = iduser
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
