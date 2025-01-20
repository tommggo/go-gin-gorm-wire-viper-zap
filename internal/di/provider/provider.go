package provider

import (
	"go-gin-gorm-wire-viper-zap/internal/cron"
	"go-gin-gorm-wire-viper-zap/internal/repository"
	"go-gin-gorm-wire-viper-zap/internal/router"
	"go-gin-gorm-wire-viper-zap/internal/service"
	"go-gin-gorm-wire-viper-zap/pkg/cache/redis"
	pkgcron "go-gin-gorm-wire-viper-zap/pkg/cron"
	"go-gin-gorm-wire-viper-zap/pkg/database/mysql"
	"go-gin-gorm-wire-viper-zap/pkg/http"

	"github.com/google/wire"
)

// InfraProvider 基础设施依赖
var InfraProvider = wire.NewSet(
	mysql.New,
	redis.New,
	http.NewServer,
)

// RepositoryProvider 数据访问层依赖
var RepositoryProvider = wire.NewSet(
	repository.NewSignalRepository,
)

// ServiceProvider 业务服务层依赖
var ServiceProvider = wire.NewSet(
	service.NewSignalService,
)

// HandlerProvider Web处理层依赖
var HandlerProvider = wire.NewSet(
	router.NewRouter,
)

// TaskProvider 任务模块依赖
var TaskProvider = wire.NewSet(
	cron.NewCronManager,
	pkgcron.New,
)

// ProviderSet 整合所有依赖
var ProviderSet = wire.NewSet(
	InfraProvider,
	RepositoryProvider,
	ServiceProvider,
	HandlerProvider,
	TaskProvider,
)
