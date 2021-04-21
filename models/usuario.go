package models

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var Usuarios map[string]*Usuario

func init() {
	Usuarios = make(map[string]*Usuario)
	// id := uuid.New().String()
	// Usuarios[id] = &Usuario{id, "Roberto", nil}
}


type Usuario struct {
	Id		string
	Nombre	string
	Contraseña	string
	Numero	string
	Email	string
	Conn	*websocket.Conn
}

func NewUsuario(nombre string, conn *websocket.Conn) Usuario {
	id := uuid.New().String()
	return Usuario{id, nombre, "", "", "", conn}
}

func (u *Usuario) Save() error {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	logs.Info(u)
	Usuarios[u.Id] = u
	return nil
}
