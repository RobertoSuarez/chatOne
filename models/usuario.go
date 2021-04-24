package models

import (
	"gorm.io/gorm"
)

var Usuarios map[string]*Usuario

func init() {
	Usuarios = make(map[string]*Usuario)
	// id := uuid.New().String()
	// Usuarios[id] = &Usuario{id, "Roberto", nil}
}

type Usuario struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Number   string `json:"number"`
	Email    string `gorm:"unique" json:"email"`
	Role	string `json:"role"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool 	`json:"isError"`
	Message string `json:"message"`
}

