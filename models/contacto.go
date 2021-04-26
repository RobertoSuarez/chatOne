package models

import (
	"gorm.io/gorm"
)

type Contacto struct {
	gorm.Model
	Number   string `json:"number"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	NickName string `json:"nickname"`
	UserID   uint   // usuario al que pertenece
}
