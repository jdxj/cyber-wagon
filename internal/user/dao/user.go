package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Nickname string
	Email    string
	Salt     string
	Password string
}

func (User) TableName() string {
	return "user"
}
