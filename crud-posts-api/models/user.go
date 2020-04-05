package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "github.com/mstfymrtc/go-posts-api/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	ID        uint       `gorm:"primary_key";json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	FullName  string     `json:"full_name"`
	UserName  string     `json:"user_name";gorm:"not null;unique"`
	Password  string     `json:"password"`
	Token     string     `json:"token";sql:"-"`
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if user.FullName == "" || user.UserName == "" || user.Password == "" {
		return u.Message(false, "Username, full name and password is required!"), false
	}
	if len(user.Password) < 6 {
		return u.Message(false, "Password length must be at least 6!"), false
	}
	temp := &User{}
	err := GetDB().Table("users").Where("user_name=?", user.UserName).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error!"), false
	}
	if temp.UserName != "" {
		return u.Message(false, "This username already exists!"), false
	}

	return u.Message(true, "User validation successful"), true
}

func (user *User) Create() (map[string]interface{}) {
	if resp, ok := user.Validate(); !ok {
		return resp
	}
	//hash password to prevent store it as plaintext
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)
	if user.ID <= 0 {
		return u.Message(false, "Failed to create user!")
	}

	//Create new JWT token for newly registered user
	tk := &Token{UserId: user.ID}
	//store user Id as claim in encrpyed token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	//sign token with token password
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Password = ""

	response := u.Message(true, "User created successfully!")
	response["user"] = user
	return response
}

func Login(userName, password string) (map[string]interface{}) {

	user := &User{}
	err := GetDB().Table("users").Where("user_name=?", userName).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User with username "+userName+" does not exist!")
		}
		return u.Message(false, "Connection error!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//password does not match
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials!")
	}
	user.Password = ""

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	//store the token in response
	user.Token = tokenString

	resp := u.Message(true, "Login successful.")
	resp["user"] = user
	return resp
}

func GetUser(u int) *User {

	user := &User{}
	GetDB().Table("users").Where("id=?", u).First(user)
	//user not found
	if user.UserName == "" {
		return nil
	}

	user.Password = ""
	return user
}
