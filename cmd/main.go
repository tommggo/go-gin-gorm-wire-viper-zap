package main

import (
	"go-gin-gorm-wire-viper-zap/internal/api"
	"go-gin-gorm-wire-viper-zap/internal/config"
	"go-gin-gorm-wire-viper-zap/internal/di"
	"go-gin-gorm-wire-viper-zap/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. 加载配置
	cfg := config.LoadConfig()

	// 2. 初始化日志
	logger.Setup(cfg)

	// 3. 初始化服务
	container, err := di.InitializeServer(cfg)
	if err != nil {
		logger.Fatal("failed to initialize server", logger.Err(err))
	}

	// 注册清理函数
	defer func() {
		// 先关闭数据库连接
		if err := container.DB.Close(); err != nil {
			logger.Error("error during database shutdown", logger.Err(err))
		}
		// 再关闭 HTTP 服务器
		if err := container.Server.Stop(); err != nil {
			logger.Error("error during server shutdown", logger.Err(err))
		}
		logger.Info("server exited")
	}()

	// 注册路由
	api.RegisterRoutes(container.Server.Engine(), container.SignalAPI)

	// 4. 优雅关闭处理
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 5. 启动服务
	serverError := make(chan error, 1)
	go func() {
		if err := container.Server.Start(); err != nil {
			serverError <- err
		}
	}()

	// 6. 等待退出信号或服务器错误
	select {
	case <-quit:
		logger.Info("shutting down server...")
	case err := <-serverError:
		logger.Fatal("server error", logger.Err(err))
	}
}
