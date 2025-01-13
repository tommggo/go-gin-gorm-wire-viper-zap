package api

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(e *gin.Engine, signalAPI *SignalAPI) {
	// 健康检查
	e.GET("/health", HealthCheck)

	// API 分组
	v1 := e.Group("/api/v1")
	{
		signal := v1.Group("/signal")
		{
			signal.POST("/create", signalAPI.Create)
			signal.GET("/:id", signalAPI.Get)
			signal.POST("/:id/process", signalAPI.Process)
		}
	}
}
