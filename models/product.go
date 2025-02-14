package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name 		string 		`gorm:"type:varchar(66);not null;index"`
	Description string 		`gorm:"type:text"`
	Type 		string 		`gorm:"type:varchar(66)"`
	CommentNum 	int 		`gorm:"type:int;default:0;column:comment_num"`
	Price 		float64 	`gorm:"type:decimal(10,2);not null"`				//表示数字共十位，小数点后两位
	IsAddCart 	bool		`gorm:"type:bool;default:false;column:is_addedCart"`
	Cover		string		`gorm:"type:text"`
	PublishTime string		`gorm:"type:varchar(20);column:publish_time"`
	Link		string		`gorm:"type:text"`
	Categories []Category	`gorm:"many2many:product_categories;"`
}

// many2many是一种多对多关系，而product_categories是中间表，用于存储product和category之间的关系
// 我们可以这么来理解， 一件衣服既可以是外套也可以是冬装，同样的,冬装既可以有衣服也可以有裤子
// 所以这就是一种多对多的关系， 我们需要一个表来存储这种关系， 这个表就是中间表

type Category struct {
	gorm.Model
	Name 		string 		`gorm:"type:varchar(66);not null;unique;index"`
	Description string 		`gorm:"type:varchar(255)"`
	Products 	[]Product 	`gorm:"many2many:product_categories;"`
}