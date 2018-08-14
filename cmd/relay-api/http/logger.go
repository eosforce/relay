package http

import (
	"time"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

// logger imp a logger for gin by seelog
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		seelog.Debugf("handler http req : %d, %v, %s, %s, %s, %s",
			statusCode, latency, clientIP, method, path, comment)
	}

}
