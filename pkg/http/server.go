package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-gin-gorm-wire-viper-zap/internal/config"
	"go-gin-gorm-wire-viper-zap/pkg/http/middleware"
	"go-gin-gorm-wire-viper-zap/pkg/logger"
)

// ErrServerClosed 服务器正常关闭的错误
var ErrServerClosed = http.ErrServerClosed

type Server struct {
	engine *gin.Engine
	srv    *http.Server
}

// RouterRegistrar 路由注册器接口
type RouterRegistrar interface {
	Register(e *gin.Engine)
}

// NewServer 创建 HTTP 服务器
func NewServer(cfg *config.Config, router RouterRegistrar) *Server {
	// 1. 设置 gin mode
	gin.SetMode(cfg.Server.Mode)

	// 2. 创建 engine
	engine := gin.New()

	// 3. 基础配置
	engine.MaxMultipartMemory = cfg.Server.MaxMultipartMemory
	engine.UseRawPath = true
	engine.UnescapePathValues = false

	// 4. 配置代理
	engine.SetTrustedProxies(nil)

	// 5. 添加中间件
	engine.Use(
		middleware.Logger(),
		middleware.ErrorHandler(),
	)

	// 6. 注册路由
	router.Register(engine)

	// 7. 创建 HTTP 服务器
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:        engine,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		MaxHeaderBytes: cfg.Server.MaxHeaderBytes,
	}

	// 8. 返回 Server
	return &Server{
		engine: engine,
		srv:    srv,
	}
}

// Engine 返回 gin 引擎，用于注册路由
func (s *Server) Engine() *gin.Engine {
	return s.engine
}

// Start 启动 HTTP 服务器
func (s *Server) Start() error {
	logger.Info("http server starting",
		logger.String("addr", s.srv.Addr),
		logger.String("mode", gin.Mode()),
	)
	return s.srv.ListenAndServe()
}

// Stop 优雅关闭 HTTP 服务器
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("shutting down http server...")
	return s.srv.Shutdown(ctx)
}
