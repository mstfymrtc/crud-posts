package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB //database

func init() {
	conn, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("failed to connect database")
	}
	db = conn
	//Migrate db
	//db.Exec("PRAGMA foreign_keys = ON")
	db.Debug().AutoMigrate(&User{}, &Post{})
	//seed static user
	Seed(db)

}
func GetDB() *gorm.DB {
	return db
}
