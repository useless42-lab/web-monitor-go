package middleware

import (
	"WebMonitor/response"
	"WebMonitor/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func CheckAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization != "" {
			userId := utils.GetUserFromToken(authorization)
			if userId != 0 {
				c.Set("userId", userId)
				c.Next()
				return
			}
		}
		response.Error(c, 9999, "未登录")
		c.Abort()
		return
	}
}
