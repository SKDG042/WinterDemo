package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID		uint	`gorm:"not null"`
	Status		string	`gorm:"type:varchar(20);not null;default:待支付"`
	TotalPrice	float64	`gorm:"not null;default:0"`
	User		User	`gorm:"foreignKey:UserID"`
}