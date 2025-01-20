package main

import (
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

	// 3. 初始化应用
	app, err := di.InitializeApp(cfg)
	if err != nil {
		logger.Fatal("failed to initialize app", logger.Err(err))
	}

	// 注册清理函数
	defer app.Stop()

	// 4. 优雅关闭处理
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 5. 启动应用
	errChan := make(chan error, 1)
	go func() {
		if err := app.Run(); err != nil {
			errChan <- err
		}
	}()

	// 6. 等待退出信号或错误
	select {
	case <-quit:
		logger.Info("shutting down gracefully...")
	case err := <-errChan:
		logger.Fatal("app error", logger.Err(err))
	}
}
