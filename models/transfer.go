package models

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	AccountOriginId      string  `gorm:"column:accountoriginid"`
	AccountDestinationId string  `gorm:"column:accountdestinationid"`
	Amount               float32 `gorm:"column:amount"`
}

type DoTransferInput struct {
	AccountDestinationId string  `json:"destination" binding:"required"`
	Amount               float32 `json:"amount" binding:"required"`
}
