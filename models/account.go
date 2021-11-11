package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string  `gorm:"column:name"`
	Cpf     string  `gorm:"unique"`
	Secret  string  `gorm:"column:secrect"`
	Balance float32 `gorm:"column:balance"`
}

type CreateAccountInput struct {
	Name   string `json:"name" binding:"required"`
	Cpf    string `json:"cpf" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}
