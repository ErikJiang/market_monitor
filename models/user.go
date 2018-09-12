package models

import (
	"github.com/jinzhu/gorm"
)



type User struct {
	gorm.Model
	UserName string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Status bool `json:"status"`
}

// 新增用户表
func AddUser(username string, password string, email string) error {
	user := User {
		UserName: username,
		Password: password,
		Email: email,
		Status: true,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 删除用户表
func DelUser(userId int) error {
	if err := db.Where("id = ?", userId).Delete(&User{}).Error; err!=nil {
		return err
	}
	return nil
}

// 更新用户表
func UpdUser(userId int, data interface{}) error {
	if err := db.Where("id = ?", userId).Update(data).Error; err!=nil {
		return err
	}
	return nil
}

// 查询用户详情
func GetUser(userId int) (user User, err error) {
	err = db.Select("id", "name", "email").Where("id = ?", userId).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return User{}, err
	}
	return user, err
}

// 查询用户列表
func GetUsers(pageNum int, pageSize int, condition interface{}) ([]User, error) {
	var users []User
	if pageNum > 0 && pageSize > 0 {
		db = db.Offset(pageNum).Limit(pageSize)
	}
	err := db.Select("id", "name", "email").Where(condition).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, err
}

