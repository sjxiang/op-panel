package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/op-panel/pkg/helper"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		
		authHeader := ctx.Request.Header.Get("Authorization")
	
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "请求 header 中 auth 为空",
			})

			ctx.Abort()
			return
		}
	
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code" : 401,  
				"msg": "请求 header 中，auth 格式有误",
			})
		
			ctx.Abort()
			return
		}	
	
		// parts[1] 是获取到的 tokenString，可以用 jwt 包方法解析
		claims, err := helper.ParseToken(parts[1]) 
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "用户认证不通过，" + err.Error(),
			})
	
			ctx.Abort()
			return
		}
	
		// 检查 expired time （重新登录，申请 JWT）
		if float64(time.Now().Unix()) > float64(claims.ExpiresAt.Unix()) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "token 过期",
			})
	
			ctx.Abort()
			return
		}
		
	
		// 将当前请求的 userID 信息保存到请求的上下文 ctx 中
		ctx.Set("user_claims", claims)
		ctx.Next()  
		
		// 后续的 handler 可以用 ctx.Get("user_claims") 来获取当前请求的用户信息
	}
}


// recovery、限流、日志