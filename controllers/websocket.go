package controllers

import (
	"chatOne/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}


type ChatController struct {
	web.Controller
}

// ejemplo localhost:8080/websocket?nombre=sdf
// @router / [get]
func (c *ChatController) Get() {
	nombre := c.GetString("nombre")

	if nombre == "" {
		c.Data["json"] = "Falta el nombre"
		c.ServeJSON()
		return
	}

	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		logs.Debug("Error: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	defer conn.Close()
	usuario := models.NewUsuario(nombre, conn)
	usuario.Save()

	for {
		metype, msg, err := conn.ReadMessage()
		if err != nil {
			logs.Error("Error: ", err)
			break
		}
		logs.Info("Mensage: ", string(msg))

		for i, _ := range models.Usuarios {
			if models.Usuarios[i].Conn == nil {
				continue
			}
			err = models.Usuarios[i].Conn.WriteMessage(metype, msg)
			if err != nil {
				logs.Error("Error: ", err)
				delete(models.Usuarios, i)
			}
		}

	}

}

// @router /users	[get]
func (c *ChatController) AllUser() {
	c.Data["json"] = models.Usuarios
	c.ServeJSON()
}
