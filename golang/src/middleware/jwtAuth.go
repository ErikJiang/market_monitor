package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/extend/code"
	"github.com/JiangInk/market_monitor/extend/jwt"
	"github.com/JiangInk/market_monitor/extend/redis"
	"github.com/rs/zerolog/log"
)

// JWTAuth Token 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 Authorization token 值
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			// 获取不到 Authorization 报：请求未携带Token,无权访问
			utils.ResponseFormat(c, code.TokenNotFound, nil)
			c.Abort()
			return
		}

		// 获取到 Token, 解析token信息
		jwtInstance := jwt.NewJWT()
		claims, err := jwtInstance.ParseToken(token)
		if err != nil {
			// 未能正常解析 Token，则报：token认证失败
			utils.ResponseFormat(c, code.TokenInvalid, nil)
			c.Abort()
			return
		}

		// 获取缓存中的Token信息
		tokenCache, err := redis.Get("TOKEN:"+claims.Email)
		if err != nil {
			log.Error().Msgf("jwt auth redis get: %v", err.Error())
			utils.ResponseFormat(c, code.ServiceInsideError, nil)
			c.Abort()
			return
		}

		// 用户注销或token失效
		if tokenCache != token {
			log.Error().Msg("user signout or token invalid")
			utils.ResponseFormat(c, code.TokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
