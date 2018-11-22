package service

import (
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/models"
	"github.com/rs/zerolog/log"
)

// UserService 用户服务层逻辑
type UserService struct{}

// QueryUserByEmail 通过邮箱查询用户信息
func (us UserService) QueryUserByEmail(email string) (user *models.User, err error) {
	userModel := &models.User{}
	condition := map[string]interface{}{
		"email": email,
	}
	user, err = userModel.FindOne(condition)
	return
}

// QueryUserByName 通过名称查询用户信息
func (us UserService) QueryUserByName(name string) (user *models.User, err error) {
	userModel := &models.User{}
	condition := map[string]interface{}{
		"user_name": name,
	}
	user, err = userModel.FindOne(condition)
	return
}

// AuthSignin 验证登录信息
func (us UserService) AuthSignin(email string, password string) (bool, error) {
	userModel := &models.User{}
	condition := map[string]interface{}{
		"email": email,
	}
	user, err := userModel.FindOne(condition)
	if err != nil {
		return false, err
	}
	if user == nil || user.Password != utils.MakeSha1(email+password) {
		return false, nil
	}
	return true, nil
}

// StoreUser 添加用户
func (us UserService) StoreUser(email string, pass string) (userID uint, err error) {
	log.Info().Msg("enter storeUser service")

	user := &models.User{
		Email:    email,
		UserName: email,
		Password: pass,
		Status:   "ENABLE",
	}
	user.Password = utils.MakeSha1(user.Email + user.Password)
	log.Debug().Msgf("user password: %s", user.Password)
	userID, err = user.Insert()
	return
}

// UpdateUser 更新用户
func (us UserService) UpdateUser(userID uint) {
	return
}

// DestroyUser 删除用户
func (us UserService) DestroyUser(userID uint) error {
	// log.Info().Msg("enter removeUser service.")
	return nil
}
