package models

import (
	"github.com/JiangInk/market_monitor/database"
	"github.com/jinzhu/gorm"
)

// User 用户表 model 定义
type User struct {
	gorm.Model
	ID       int64  `json:"id" gorm:"primary_key"`
	UserName string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
}

// Insert 新增用户
func (user *User) Insert() (id int64, err error) {

	result := database.DB.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

// FindOne 查询用户详情
func (user *User) FindOne(userID int64) (userInfo User, err error) {
	err = database.DB.Select("id", "name", "email").Where("id = ?", userID).First(&userInfo).Error
	if err != nil {
		return
	}
	return
}

// FindAll 获取用户列表
func (user *User) FindAll(pageNum int, pageSize int, condition interface{}) (users []User, err error) {

	result := database.DB.Offset(pageNum).Limit(pageSize).Select("id", "name", "email").Where(condition).Find(&users)
	err = result.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

// UpdateOne 修改用户
func (user *User) UpdateOne(userID int64) (updUser User, err error) {
	if err = database.DB.Select([]string{"id", "username"}).First(&updUser, userID).Error; err != nil {
		return
	}

	//参数1:需要修改的源用户
	//参数2:修改更新的数据
	if err = database.DB.Model(&updUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

// DeleteOne 删除用户
func (user *User) DeleteOne(id int64) (delUser User, err error) {
	if err = database.DB.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = database.DB.Delete(&user).Error; err != nil {
		return
	}
	delUser = *user
	return
}
