package middleware

/********************************************************************************
* Temancode Example Add Middleware Package                                      *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import "github.com/gin-gonic/gin"

func AddDefaultHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Content-Security-Policy", "default-src 'self'")
	}
}
