package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)


type ContactController struct {
	web.Controller
}

// @router / [get]
func (c *ContactController) Get() {
	// envia los contato en json
	iduser := c.GetString(":iduser")
	idcontact := c.GetString(":idcontact")
	c.Ctx.WriteString(fmt.Sprintln("iduser: ", iduser, "idcontact: ", idcontact))
}


func (c *ContactController) AllContact() {
	c.Data["json"] = "Todos los contactos de " + c.GetString(":iduser")
	c.ServeJSON()
}
