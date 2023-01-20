package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

var sugar = zap.NewExample().Sugar()

// HeaderCheck - is a checking Header X-PING for the ping value for installing Header X-PONG with value pong :D
func HeaderCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-PING") == "ping" {
			c.Writer.Header().Set("X-PONG", "pong")
			sugar.Info("Received X-PING header with value ping, sent X-PONG header with value pong", " time of:", time.Now().Format("02.01.2006.15.04.05"))
		} else if c.Request.Header.Get("X-PING") == "" {
		} else {
			sugar.Info("The X-PING header was specified incorrectly", " time of:", time.Now().Format("02.01.2006.15.04.05"))
		}
	}
}
