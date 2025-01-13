package api

import (
	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	Success(c, map[string]string{
		"status":  "ok",
		"message": "service is running",
	})
}
