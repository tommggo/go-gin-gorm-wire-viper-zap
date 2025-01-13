package middleware

import (
	"fmt"
	"go-gin-gorm-wire-viper-zap/internal/api"
	"go-gin-gorm-wire-viper-zap/internal/errors"
	"go-gin-gorm-wire-viper-zap/pkg/logger"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// 处理 panic
			if err := recover(); err != nil {
				// 记录错误日志
				logger.Error("panic recovered",
					logger.Any("error", err),
					logger.String("path", c.Request.URL.Path),
				)

				// 包装为系统错误
				appErr := errors.Wrap(errors.ErrSystem, fmt.Errorf("%v", err))
				api.Error(c, appErr)
				c.Abort()
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
