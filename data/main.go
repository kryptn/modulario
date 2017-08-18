package data

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"time"
)

type Engine struct {
	db      *gorm.DB
	Dialect string
	Args    string
	LogMode bool
}

func (e *Engine) InitDB() {
	var err error
	e.db, err = gorm.Open(e.Dialect, e.Args)
	if err != nil {
		log.Fatalf("Database error when connecting: '%v'", err)
	}
	e.db.LogMode(e.LogMode)
}

func (e *Engine) InitSchema() {
	rand.Seed(time.Now().UTC().UnixNano())
	e.db.AutoMigrate(&User{}, &Post{}, &Link{})
	e.db.Model(&Link{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "RESTRICT")
}
