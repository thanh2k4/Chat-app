package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		log.Printf("[%s] %s %s - %v",
			c.Request.Method, c.Request.URL.Path, c.Request.UserAgent(), time.Since(startTime))
	}
}
