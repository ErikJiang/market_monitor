package jwt

import (
	"time"
	"github.com/JiangInk/market_monitor/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/JiangInk/market_monitor/config"
)

// JWT 认证相关
type JWT struct {
	JWTSecret []byte
}

// NewJWT 创建 JWT 实例
func NewJWT() *JWT {
	return &JWT {
		[]byte(config.ServerConf.JWTSecret),
	}
}

// CustomClaims jwt信息
type CustomClaims struct {
	ID			string	`json:userId`
	UserName	string	`json:"username"`
	Email		string	`json:"email"`
	jwt.StandardClaims
}

// CreateToken 生成 Token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(j.JWTSecret)
}

// ParseToken 解析 Token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JWTSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
