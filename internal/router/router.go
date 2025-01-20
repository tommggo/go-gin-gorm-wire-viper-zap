package router

import (
	"go-gin-gorm-wire-viper-zap/internal/api"
	"go-gin-gorm-wire-viper-zap/internal/service"
	"go-gin-gorm-wire-viper-zap/pkg/http"

	"github.com/gin-gonic/gin"
)

// 实现 http.RouterRegistrar 接口
type Router struct {
	signalService service.SignalService
}

func NewRouter(signalService service.SignalService) http.RouterRegistrar {
	return &Router{
		signalService: signalService,
	}
}

// Register 注册所有路由
func (r *Router) Register(e *gin.Engine) {
	// 健康检查
	r.registerHealthRoutes(e)

	// API 路由
	v1 := e.Group("/api/v1")
	r.registerSignalRoutes(v1)
}

func (r *Router) registerHealthRoutes(e *gin.Engine) {
	e.GET("/health", HealthCheck)
}

func (r *Router) registerSignalRoutes(v1 *gin.RouterGroup) {
	signalAPI := api.NewSignalAPI(r.signalService)

	signal := v1.Group("/signal")
	signal.POST("/create", signalAPI.Create)
	signal.GET("/:id", signalAPI.Get)
	signal.POST("/:id/process", signalAPI.Process)
}
