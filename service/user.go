package service

import (
	model "github.com/JiangInk/market_monitor/models"
	helper "github.com/JiangInk/market_monitor/helpers"
    "github.com/rs/zerolog/log"
)

type UserSerivce struct{}


var validate *validator.Validate

// 添加用户
func (us UserSerivce) addUser(email string, pass: string) (err error) {
	log.Info().Msg("enter signup service.")

	user := &model.User{
		Email:	email,
		UserName: email,
		Password: pass,
		Status: true,
	}
	if len(user.Email) == 0 || len(user.Password) == 0 {
		err = errors.New("email or password cannot be nil.")
	} else {
		user.Password = helper.Md5(user.Email+user.Password)
		err = user.Insert()
	}
}

// 移除用户
func (us UserSerivce) removeUser(userId int) error {
	log.Info().Msg("enter removeUser service.")
	validate = validator.New()
	err := validate.Var(userId, "required")
	if err != nil {
		log.Error().Msgf("%v", err)
		return err
	}

	// 更新状态为禁用false
	return models.UpdUser(userId, map[string]bool {"status": false})
}

// 修改用户密码
func (us UserSerivce) alterPasswd(userId int, passwd string) error {
	log.Info().Msg("enter alterPasswd service.")

	validate = validator.New()
	err := validate.Var(userId, "required")
	if err != nil {
		log.Error().Msgf("%v", err)
		return err
	}
	err = validate.Var(passwd, "required")
	if err != nil {
		log.Error().Msgf("%v", err)
		return err
	}

	return models.UpdUser(userId, map[string]string { "password": passwd })
}

// 修改用户邮箱
func (us UserSerivce) alterEmail(userId int, email string) error {
	log.Info().Msg("enter alterEmail service.")

	validate = validator.New()
	err := validate.Var(userId, "required")
	if err != nil {
		log.Error().Msgf("%v", err)
		return err
	}
	err = validate.Var(email, "required,email")
	if err != nil {
		log.Error().Msgf("%v", err)
		return err
	}

	return models.UpdUser(userId, map[string]string { "password": email })
}

// 查询用户详情
func (us UserSerivce) getDetail(userId int) (user models.User, err error) {
	log.Info().Msg("enter getDetail service.")

	validate = validator.New()
	err = validate.Var(userId, "required")
	if err != nil {
		log.Error().Msgf("%v", err)
		return models.User{}, err
	}
	return models.GetUser(userId)
}

