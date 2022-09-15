package dao

import "gorm.io/gorm"

type File struct {
	gorm.Model

	UserID   uint64
	Filename string
	MD5      string
}

func (File) TableName() string {
	return "file"
}