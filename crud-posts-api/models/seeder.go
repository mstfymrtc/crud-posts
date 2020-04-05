package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var user = User{
	UserName: "mustafayumurtaci",
	FullName: "Mustafa Yumurtacı",
	Password: "123qwe",
}

func Seed(db *gorm.DB) {

	count := 0
	err := db.Table("users").Count(&count).Error
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		err := db.Debug().Model(&User{}).Create(&user).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

}
