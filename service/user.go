package service

import (
	"github.com/JiangInk/market_monitor/models"
	"gopkg.in/go-playground/validator.v9"

    "github.com/rs/zerolog/log"
)

type User struct {
	UserName	string	`validator:"required"`
	Password	string	`validator:"required"`
	Email		string	`validator:"required,email"`
}

var validate *validator.Validate

// 创建用户
func createUser(userInfo User) error {
	log.Info().Msg("enter signup service.")
	validate = validator.New()
	err := validate.Struct(userInfo)
	if err != nil {
		log.Error().Msgf("%v", err)
		return err
	}

	return models.AddUser(userInfo.UserName, userInfo.Password, userInfo.Email)
}

// 移除用户
func removeUser(userId int) error {
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
func alterPasswd(userId int, passwd string) error {
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
func alterEmail(userId int, email string) error {
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
func getDetail(userId int) (user models.User, err error) {
	log.Info().Msg("enter getDetail service.")

	validate = validator.New()
	err = validate.Var(userId, "required")
	if err != nil {
		log.Error().Msgf("%v", err)
		return models.User{}, err
	}
	return models.GetUser(userId)
}

