package models

import (
	"deneme/pgk/config"
	_ "deneme/pgk/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model

	FirstName string `gorm:""json:"firstName"`
	LastName  string `gorm:""json:"lastName"`
	Email     string `gorm:""json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)

	return u

}

func GetAllUser() []User {

	var Users []User
	db.Find(&Users)

	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {

	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)

	return &getUser, db

}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(user)

	return user
}
