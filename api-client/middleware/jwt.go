package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/util"

	"net/http"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		token := ctx.Query("token")
		if token == "" {
			code = http.StatusUnauthorized
			data = gin.H{
				"code": code,
				"msg":  http.StatusText(code),
			}
			ctx.JSON(code, data)
			ctx.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
				data = gin.H{
					"code": code,
					"msg":  http.StatusText(code),
				}
				ctx.JSON(code, data)
				ctx.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusUnauthorized
				data = gin.H{
					"code": code,
					"msg":  http.StatusText(code),
				}
				ctx.JSON(code, data)
				ctx.Abort()
				return
			}
			ctx.Set("user_id", claims.UserId)
			ctx.Set("username", claims.Username)

			ctx.Next()
		}

	}
}
