package middleware

import "github.com/gin-gonic/gin"

func AddDefaultHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Content-Security-Policy", "default-src 'self'")
	}
}
