package db

import (
	"gin-gorm-boilerplate/config"
	repo "gin-gorm-boilerplate/db/repositories"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Init database from config with postgres
func Init() {
	db = connectDB()
	defer db.Close()
	db.AutoMigrate(
		&repo.User{},
	)
}

func connectDB() *gorm.DB {
	config := config.GetConfig()
	host := config.GetString("db.host")
	port := config.GetString("db.port")
	username := config.GetString("db.username")
	password := config.GetString("db.password")
	dbname := config.GetString("db.database")
	sslmode := config.GetString("db.sslmode")
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+dbname+" password="+password+" sslmode="+sslmode)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// GetDB returns db object
func GetDB() *gorm.DB {
	if err := db.DB().Ping(); err != nil {
		db = connectDB()
	}
	return db
}
