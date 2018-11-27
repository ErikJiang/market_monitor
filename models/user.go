package models

import (
	"github.com/jinzhu/gorm"
)

// User 用户表 model 定义
type User struct {
	gorm.Model
	UserName	string	`gorm:"unique_index;default:null"`
	Password	string	`gorm:"default:null"`
	Email		string	`gorm:"unique_index;default:null"`
	Avatar		string	`gorm:"default:null"`
	Status		string	`sql:"type:ENUM('ENABLE', 'DISABLE')"`
}

// Insert 新增用户
func (user *User) Insert() (userID uint, err error) {

	result := DB.Create(&user)
	userID = user.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

// FindOne 查询用户详情
func (user *User) FindOne(condition map[string]interface{}) (*User, error) {
	var userInfo User
	result := DB.Select("id, user_name, email, password").Where(condition).First(&userInfo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if userInfo.ID > 0 {
		return &userInfo, nil
	}
	return nil, nil
}

// FindAll 获取用户列表
func (user *User) FindAll(pageNum int, pageSize int, condition interface{}) (users []User, err error) {

	result := DB.Offset(pageNum).Limit(pageSize).Select("id", "name", "email").Where(condition).Find(&users)
	err = result.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

// UpdateOne 修改用户
func (user *User) UpdateOne(userID uint, data map[string]interface{}) (*User, error) {
	err := DB.Model(&User{}).Where("id = ?", userID).Updates(data).Error
	if err != nil {
		return nil, err
	}
	var updUser User
	err = DB.Select([]string{"id", "user_name", "email"}).First(&updUser, userID).Error
	if err != nil {
		return nil, err
	}
	return &updUser, nil
}

// DeleteOne 删除用户
func (user *User) DeleteOne(userID uint) (delUser User, err error) {
	if err = DB.Select([]string{"id"}).First(&user, userID).Error; err != nil {
		return
	}

	if err = DB.Delete(&user).Error; err != nil {
		return
	}
	delUser = *user
	return
}
