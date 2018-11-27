package service

import (
	"github.com/JiangInk/market_monitor/extend/code"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/models"
	"github.com/rs/zerolog/log"
)

// UserService 用户服务层逻辑
type UserService struct{
	UserID		uint
	Email		string
	Name		string
	Password	string
}

// QueryUserByEmail 通过邮箱查询用户信息
func (us *UserService) QueryUserByEmail(email string) (user *models.User, err error) {
	userModel := &models.User{}
	condition := map[string]interface{}{
		"email": email,
	}
	user, err = userModel.FindOne(condition)
	return
}

// QueryUserByName 通过名称查询用户信息
func (us *UserService) QueryUserByName(name string) (user *models.User, err error) {
	userModel := &models.User{}
	condition := map[string]interface{}{
		"user_name": name,
	}
	user, err = userModel.FindOne(condition)
	return
}

// AuthSignin 验证登录信息
func (us *UserService) AuthSignin(email string, password string) (bool, error) {
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
func (us *UserService) StoreUser(email string, pass string) (userID uint, err error) {
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

// UpdateUserName 修改用户昵称
func (us *UserService) UpdateUserName(userName string) (*models.User, *code.Code) {
	userModel := &models.User{}
	// 查询用户名是否已被使用
	user, err := userModel.FindOne(map[string]interface{}{"user_name": userName})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, code.ServiceInsideError
	}

	if user != nil {
		return nil, code.AccountNameExist
	}

	// 更新用户名称
	updateUser, err := userModel.UpdateOne(us.UserID, map[string]interface{}{
		"user_name": userName,
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, code.ServiceInsideError
	}
	return updateUser, nil
}

// UpdateUserPass 修改用户密码
func (us *UserService) UpdateUserPass(oldPass, newPass string) (*models.User, *code.Code) {
	userModel := &models.User{}
	// 1. 获取对应用户信息（密码）
	user, err := userModel.FindOne(map[string]interface{}{"email": us.Email})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, code.ServiceInsideError
	}

	// 2. 为参数中原密码明文做hash,然后与用户密码hash进行校验对比
	oldPassHash := utils.MakeSha1(us.Email+oldPass)
	if oldPassHash != user.Password {
		return nil, code.AccountPassUnmatch
	}
	// 3. 校验通过，则更新用户密码
	updateUser, err := userModel.UpdateOne(user.ID, map[string]interface{}{
		"password": utils.MakeSha1(us.Email+newPass),
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, code.ServiceInsideError
	}
	return updateUser, nil
}

// UpdateUserAvatar 修改用户头像
func (us *UserService) UpdateUserAvatar(avatar string) (*models.User, *code.Code) {
	// 3. 校验通过，则更新用户密码
	userModel := &models.User{}
	updateUser, err := userModel.UpdateOne(us.UserID, map[string]interface{}{
		"avatar": avatar,
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, code.ServiceInsideError
	}
	return updateUser, nil
}


// DestroyUser 删除用户
func (us *UserService) DestroyUser(userID uint) error {
	// log.Info().Msg("enter removeUser service.")
	return nil
}
