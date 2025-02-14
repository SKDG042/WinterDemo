package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username 		string `gorm:"type:varchar(20);unique;not null;index"`
	Password 		string `gorm:"type:varchar(255);not null"`
	Nickname 		string `gorm:"type:varchar(20)"`
	Avatar 			string `gorm:"type:text"`
	Introduction 	string `gorm:"type:text"`
	Telephone    	string `gorm:"type:varchar(11)"`
	QQ				string `gorm:"type:varchar(10)"`
	Gender			string `gorm:"type:varchar(10)"`
	Email			string `gorm:"type:varchar(50)"`
	Birthday		string `gorm:"type:varchar(10)"`
}


