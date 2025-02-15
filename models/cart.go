package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID		uint		`gorm:"type:uint;not null;index;column:user_id"`
	ProductID	uint		`gorm:"type:uint;not null;index;column:product_id"`
	Quantity	int			`gorm:"type:int;not null;default:0"`
	// Fix 一个购物车可以有多个商品，一个商品也可以在多个购物车中
	Product 	[]Product	`gorm:"many2many:cart_products;"`
}