package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"go-gin-gorm-wire-viper-zap/pkg/logger"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		latency := end.Sub(start)

		// 获取状态
		status := c.Writer.Status()

		// 记录日志，格式类似 Gin 默认的访问日志
		logger.Info("[GIN]",
			logger.String("method", c.Request.Method),
			logger.Int("status", status),
			logger.Duration("latency_s", latency),
			logger.String("ip", c.ClientIP()),
			logger.String("path", path),
			logger.String("query", query),
			logger.Int("size", c.Writer.Size()),
		)
	}
}
