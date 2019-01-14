package jwt

import (
	"errors"
	"time"

	"github.com/ErikJiang/market_monitor/extend/conf"
	"github.com/dgrijalva/jwt-go"
)

// JWT 认证相关
type JWT struct {
	JWTSecret []byte
}

// NewJWT 创建 JWT 实例
func NewJWT() *JWT {
	return &JWT{[]byte(conf.ServerConf.JWTSecret)}
}

var (
	// ErrTokenExpired 验证令牌失效
	ErrTokenExpired = errors.New("Token is expired")
	// ErrTokenNotValidYet 验证令牌未激活
	ErrTokenNotValidYet = errors.New("Token not active yet")
	// ErrTokenMalformed 验证并非属于令牌
	ErrTokenMalformed = errors.New("That's not even a token")
	// ErrTokenInvalid 验证为无效的令牌
	ErrTokenInvalid = errors.New("Couldn't handle this token")
)

// CustomClaims jwt信息
type CustomClaims struct {
	ID       uint   `json:"userId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
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
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 刷新 Token
func (j *JWT) RefreshToken(token string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JWTSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		jwt.TimeFunc = time.Now
		expiredTime := time.Duration(conf.ServerConf.JWTExpire)
		claims.StandardClaims.ExpiresAt = time.Now().Add(expiredTime * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", ErrTokenInvalid
}
