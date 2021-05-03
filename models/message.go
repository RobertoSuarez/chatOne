package models

import "gorm.io/gorm"

// Informacion que se envia al cliente
type Info struct {
	Data    interface{} `json:"data"`
	IsError bool        `json:"iserror"`
	Message string      `json:"message"`
}

func NewInfo() *Info {
	return &Info{}
}

func (m *Info) OK(msg string) {
	m.IsError = false
	m.Message = msg
}

func (m *Info) SetInfo(data interface{}, iserror bool, msg string) {
	m.Data = data
	m.IsError = iserror
	m.Message = msg
}

// Mensajes de comunicación con usuarios
type Message struct {
	gorm.Model
	Text     string `json:"text"`
	IDUser   uint   `json:"iduser"`
	IDConver uint   `json:"idconver"`
}
