package service

import (
	"time"

	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/extend/jwt"
	"github.com/JiangInk/market_monitor/extend/redis"
	"github.com/JiangInk/market_monitor/models"
	goJWT "github.com/dgrijalva/jwt-go"
)

// AuthService 认证相关
type AuthService struct {
	User *models.User
}

// GenerateToken 生成 Token
func (as *AuthService) GenerateToken(user models.User) (string, error) {
	jwtInstance := jwt.NewJWT()
	nowTime := time.Now()
	expireTime := time.Duration(config.ServerConf.JWTExpire)
	claims := jwt.CustomClaims{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: goJWT.StandardClaims{
			ExpiresAt: nowTime.Add(expireTime * time.Hour).Unix(),
			Issuer:    "monitor",
		},
	}
	// 创建token
	token, err := jwtInstance.CreateToken(claims)
	if err != nil {
		return "", err
	}

	// 设置redis缓存
	const hourSecs int = 60 * 60
	redis.Set("TOKEN:"+user.Email, token, config.ServerConf.JWTExpire * hourSecs)
	return token, nil
}

// DestroyToken 销毁 Token
func (as *AuthService) DestroyToken(email string) (bool, error) {
	return redis.Del("TOKEN:"+email)
}
