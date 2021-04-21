package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var Usuarios map[string]*Usuario

func init() {
	Usuarios = make(map[string]*Usuario)
	id := uuid.New().String()
	Usuarios[id] = &Usuario{id, "Roberto", nil}
}


type Usuario struct {
	Id	string
	Nombre	string
	Conn	*websocket.Conn
}

func NewUsuario(nombre string, conn *websocket.Conn) Usuario {
	id := uuid.New().String()
	return Usuario{id, nombre, conn}
}

func (u *Usuario) Save() error {
	if u.Id == "" {
		return errors.New("No tiene id")
	}
	Usuarios[u.Id] = u
	return nil
}
