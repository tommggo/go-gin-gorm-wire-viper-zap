package router

import (
	"go-gin-gorm-wire-viper-zap/internal/api"

	"github.com/gin-gonic/gin"
)

type Router struct {
	signalAPI *api.SignalAPI
}

func NewRouter(signalAPI *api.SignalAPI) *Router {
	return &Router{
		signalAPI: signalAPI,
	}
}

// Register 注册所有路由
func (r *Router) Register(e *gin.Engine) {
	// 健康检查
	e.GET("/health", HealthCheck)

	// API 路由
	v1 := e.Group("/api/v1")
	r.registerSignalRoutes(v1)
}

func (r *Router) registerSignalRoutes(v1 *gin.RouterGroup) {
	signal := v1.Group("/signal")
	signal.POST("/create", r.signalAPI.Create)
	signal.GET("/:id", r.signalAPI.Get)
	signal.POST("/:id/process", r.signalAPI.Process)
}
