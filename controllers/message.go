package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MessageController struct {
	web.Controller
}

// @Param	iduser	path	int	true	"id del usuario"
// @router /:iduser/message [get]
func (m *MessageController) GetAll() {
	iduser := m.GetString(":iduser")
	m.Data["json"] = "Todos los mensajes de: " + iduser
	m.ServeJSON()
}

//client have send application/json
/*
{
	"recipient": {
		"id": "iduser"
	},
	"message": {
    	"text": "Hola"
   	}
}
*/

// @Param	iduser	path	int	true	"id del usuario"
// @router /:iduser/message [post]
func (m *MessageController) Post() {
	iduser := m.GetString(":iduser")
	m.Data["json"] = "Mensaje enviado " + iduser
	m.ServeJSON()
}

// OneMessage retorna el mensaje que este asociado al id
// @Param	iduser	path	int	true	"id del usuario"
// @Param	idmessage	path	int	true	"Identificación del message"
// @router /:iduser/message/:idmessage [get]
func (m *MessageController) OneMessage() {
	iduser := m.GetString(":iduser")
	idmsg := m.GetString(":idmessage")

	m.Data["json"] = "Message localizado " + iduser + " - " + idmsg
	m.ServeJSON()
}
