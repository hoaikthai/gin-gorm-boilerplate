package repositories

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	UserName string `gorm:"type:varchar(20);unique_index"`
	Password string `gorm:"type:varchar(20)"`
}
