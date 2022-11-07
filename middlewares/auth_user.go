package middlewares

import (
	"gin-gorm-oj/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthUserCheck
// 验证用户是否是管理员的中间件
func AuthUserCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Check if user is admin
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
		if userClaim == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
		c.Next()
		c.Set("user", userClaim)
	}
}
