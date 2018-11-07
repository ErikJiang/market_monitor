package service

import (
	"time"
	"github.com/JiangInk/market_monitor/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/JiangInk/market_monitor/config"
)

var jwtSecret = []byte(config.ServerConf.JWTSecret)

// Claims jwt信息
type Claims struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// AuthService 认证相关
type AuthService struct{}

// MakeToken 生成 Token
func (as AuthService) MakeToken(email string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	userModel := &models.User{}
	condition := map[string]interface{}{
		"email": email,
	}

	user, err := userModel.FindOne(condition)
	if err != nil {
		return "", err
	}
	claims := Claims{
		user.UserName,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "market_monitor",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析 Token
func (as AuthService) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
