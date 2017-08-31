package http

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

	// conditional things

	DeciderType  string
	Conditionals []Conditional
	Decider      func() Link `gorm:"-" json:"-"`
}

type Link struct {
	gorm.Model
	Url            string
	LastHttpStatus uint `gorm:"default:200"`
	Accesses       uint `gorm:"default:0"`
	PostID         uint `gorm:"index"`

	// this is for conditional stuffs
	Condition Conditional
}

type AttachedPost struct {
	gorm.Model
	ReferredPostID string
	PostID         uint `gorm:"index"`

	Metric uint

	Post Post
}

type Conditional struct {
	gorm.Model

	Threshold uint

	LinkID uint `gorm:"index"`
	PostID uint `gorm:"index"`
}
