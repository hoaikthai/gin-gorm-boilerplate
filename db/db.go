package db

import (
	"gin-gorm-boilerplate/config"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Init database from config with postgres
func Init() {
	config := config.GetConfig()
	host := config.GetString("db.host")
	port := config.GetString("db.port")
	username := config.GetString("db.username")
	password := config.GetString("db.password")
	dbname := config.GetString("db.dbname")
	db, err := gorm.Open("postgres", "host="+host+"port="+port+"user="+username+"dbname="+dbname+"password="+password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

// GetDB returns db object
func GetDB() *gorm.DB {
	return db
}
