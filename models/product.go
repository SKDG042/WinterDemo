package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name 		string 		`gorm:"type:varchar(66);not null"`
	Description string 		`gorm:"type:text"`
	Type 		string 		`gorm:"type:varchar(66)"`
	CommentNum 	int 		`gorm:"type:int;default:0;column:comment_num"`
	Price 		float64 	`gorm:"type:decimal(10,2);not null"`//共十位，小数点后两位
	IsAddCart 	bool		`gorm:"type:bool;default:false;column:is_addedCart"`
	Cover		string		`gorm:"type:varchar(255)"`
}

type Category struct {
	gorm.Model
	Name 		string 		`gorm:"type:varchar(66);not null"`
	Description 	string 		`gorm:"type:varchar(255)"`
	Products 	[]Product 	`gorm:"gorm`
}