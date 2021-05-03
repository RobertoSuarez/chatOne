package models

import (

	"gorm.io/gorm"
)

type Conversación struct {
	gorm.Model
	IDUser1 uint `json:"iduser1"`
	IDUser2 uint `json:"iduser2"`
}
