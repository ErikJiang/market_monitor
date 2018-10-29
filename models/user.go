package models

import (
	orm "github.com/JiangInk/market_monitor/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
}

// 新增用户
func (user *User) Insert() (id int64, err error) {
	result := orm.db.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

// 查询用户详情
func (user *User) FindOne(userId int64) (user User, err error) {
	if err = orm.db.Select("id", "name", "email").Where("id = ?", userId).First(&user).Error; err != nil {
		return
	}
	return
}

// 获取用户列表
func (user *User) FindAll(pageNum int, pageSize int, condition interface{}) (users []User, err error) {

	if pageNum > 0 && pageSize > 0 {
		db = orm.db.Offset(pageNum).Limit(pageSize)
	}
	err := db.Select("id", "name", "email").Where(condition).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

// 修改用户
func (user *User) UpdateOne() (updUser User, err error) {
	if err = orm.db.Select([]string{"id", "username"}).First(&updUser, id).Error; err != nil {
		return
	}

	//参数1:需要修改的源用户
	//参数2:修改更新的数据
	if err = orm.db.Model(&updUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

// 删除用户
func (user *User) DeleteOne(id int64) (delUser User, err error) {
	if err = orm.db.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.db.Delete(&user).Error; err != nil {
		return
	}
	delUser = *user
	return
}
