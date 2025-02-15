package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content 	string		`gorm:"type:text;not null"`
	UserID		uint		`gorm:"type:uint;not null;index;column:user_id"`
	Avatar		string		`gorm:"type:text"`
	Nickname 	string		`gorm:"type:varchar(20);not null"`
	PraiseNum 	int			`gorm:"type:int;default:0;column:praise_num"`
	IsPraise	uint		`gorm:"type:uint;default:0;column:is_praise"` // 0为未处理， 1为点赞， 2为点踩
	ProductID	uint		`gorm:"type:uint;not null;index;column:product_id"`
	ParentID	uint		`gorm:"type:uint;default:0;column:parent_id"` // 父评论ID, 0为顶级评论
	// gorm:"-"的意思是这个字段不会存储在数据库中，以防止数据冗余等等
	Children    []Comment	`gorm:"-"`
}