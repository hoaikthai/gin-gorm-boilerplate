package auth

import (
	"gin-gorm-boilerplate/db"
	"gin-gorm-boilerplate/db/repositories"
)

func isExist(f SignUpForm) bool {
	db := db.GetDB()
	defer db.Close()
	user := repositories.User{}
	db.Where("email = ? OR user_name = ?", f.Email, f.UserName).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func createUser(f SignUpForm) error {
	db := db.GetDB()
	defer db.Close()
	f.Password = hashPassword(f.Password)
	user := repositories.User{Email: f.Email, UserName: f.UserName, Password: f.Password}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
