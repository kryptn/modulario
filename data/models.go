package data

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	Password string
	Posts    []Post // one to many
}

type Post struct {
	gorm.Model
	Key    string `gorm:"index"`
	UserID uint   `gorm:"index"`
	Links  []Link // one to many
}

type Link struct {
	gorm.Model
	Url            string
	LastHttpStatus uint `gorm:"default:200"`
	Accesses       uint `gorm:"default:0"`
	PostID         uint `gorm:"index"`
}
