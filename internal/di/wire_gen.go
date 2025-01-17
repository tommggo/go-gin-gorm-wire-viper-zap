// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"go-gin-gorm-wire-viper-zap/internal/config"
	"go-gin-gorm-wire-viper-zap/internal/cron"
	"go-gin-gorm-wire-viper-zap/internal/repository"
	"go-gin-gorm-wire-viper-zap/internal/router"
	"go-gin-gorm-wire-viper-zap/internal/service"
	"go-gin-gorm-wire-viper-zap/pkg/cache"
	"go-gin-gorm-wire-viper-zap/pkg/cache/redis"
	cron2 "go-gin-gorm-wire-viper-zap/pkg/cron"
	"go-gin-gorm-wire-viper-zap/pkg/database"
	"go-gin-gorm-wire-viper-zap/pkg/database/mysql"
	"go-gin-gorm-wire-viper-zap/pkg/http"
	"go-gin-gorm-wire-viper-zap/pkg/logger"
)

// Injectors from wire.go:

// InitializeApp 初始化应用
func InitializeApp(cfg *config.Config) (*App, error) {
	db, err := mysql.New(cfg)
	if err != nil {
		return nil, err
	}
	signalRepository := repository.NewSignalRepository(db)
	signalService := service.NewSignalService(signalRepository)
	routerRegistrar := router.NewRouter(signalService)
	server := http.NewServer(cfg, routerRegistrar)
	cache, err := redis.New(cfg)
	if err != nil {
		return nil, err
	}
	registrar := cron.NewCronManager(signalService, cfg)
	cronCron := cron2.New(cache, registrar)
	app := &App{
		server: server,
		db:     db,
		cache:  cache,
		cron:   cronCron,
	}
	return app, nil
}

// wire.go:

// App 应用程序结构
type App struct {
	server *http.Server
	db     database.DB
	cache  cache.Cache
	cron   *cron2.Cron
}

// Run 启动应用
func (a *App) Run() error {

	a.cron.Start()

	return a.server.Start()
}

// Stop 优雅关闭
func (a *App) Stop() {

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
