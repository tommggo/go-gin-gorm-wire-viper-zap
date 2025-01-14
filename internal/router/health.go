package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "service is running",
	})
}
