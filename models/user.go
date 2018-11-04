package models

import (
	"github.com/jinzhu/gorm"
)

// User 用户表 model 定义
type User struct {
	gorm.Model
	UserName string `gorm:"unique_index;default:null"`
	Password string `gorm:"default:null"`
	Email    string `gorm:"unique_index;default:null"`
	Status   string `sql:"type:ENUM('ENABLE', 'DISABLE')"`
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
func (user *User) FindOne(condition map[string]interface{}) (userInfo User, err error) {
	result := DB.Select("id, user_name, email").Where(condition).First(&userInfo)
	err = result.Error
	return
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
func (user *User) UpdateOne(userID uint) (updUser User, err error) {
	if err = DB.Select([]string{"id", "username"}).First(&updUser, userID).Error; err != nil {
		return
	}

	//参数1:需要修改的源用户
	//参数2:修改更新的数据
	if err = DB.Model(&updUser).Updates(&user).Error; err != nil {
		return
	}
	return
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
