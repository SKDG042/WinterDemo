package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID		uint	`gorm:"type:uint;not null;index;column:user_id"`
	ProductID	uint	`gorm:"type:uint;not null;index;column:product_id"`
	Quantity	int		`gorm:"type:int;not null;default:0"`
	Product 	Product	`gorm:"foreignKey:ProductID"`
}