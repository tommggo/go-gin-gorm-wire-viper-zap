//go:build wireinject
// +build wireinject

package di

import (
	"go-gin-gorm-wire-viper-zap/internal/config"
	"go-gin-gorm-wire-viper-zap/internal/di/provider"
	"go-gin-gorm-wire-viper-zap/pkg/cache"
	"go-gin-gorm-wire-viper-zap/pkg/cron"
	"go-gin-gorm-wire-viper-zap/pkg/database"
	"go-gin-gorm-wire-viper-zap/pkg/http"
	"go-gin-gorm-wire-viper-zap/pkg/logger"

	"github.com/google/wire"
)

// App 应用程序结构
type App struct {
	server *http.Server
	db     database.DB
	cache  cache.Cache
	cron   *cron.Cron
}

// Run 启动应用
func (a *App) Run() error {
	// 启动定时任务
	a.cron.Start()

	// 启动 HTTP 服务
	return a.server.Start()
}

// Stop 优雅关闭
func (a *App) Stop() {
	// 按依赖顺序关闭组件
	a.cron.Stop()
	if err := a.cache.Close(); err != nil {
		logger.Error("close cache failed", logger.Err(err))
	}
	if err := a.db.Close(); err != nil {
		logger.Error("close db failed", logger.Err(err))
	}
	if err := a.server.Stop(); err != nil {
		logger.Error("close server failed", logger.Err(err))
	}
	logger.Info("app stopped")
}

// InitializeApp 初始化应用
func InitializeApp(cfg *config.Config) (*App, error) {
	wire.Build(
		provider.ProviderSet,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
