package provider

import (
	"go-gin-gorm-wire-viper-zap/internal/api"
	"go-gin-gorm-wire-viper-zap/internal/repository"
	"go-gin-gorm-wire-viper-zap/internal/service"
	"go-gin-gorm-wire-viper-zap/pkg/database"
	"go-gin-gorm-wire-viper-zap/pkg/http"

	"github.com/google/wire"
)

// DBProvider 数据库相关依赖
var DBProvider = wire.NewSet(
	database.NewDB,
)

// ServerProvider HTTP 服务器依赖
var ServerProvider = wire.NewSet(
	http.NewServer,
)

// RepositoryProvider Repository 层依赖
var RepositoryProvider = wire.NewSet(
	repository.NewSignalRepository,
	wire.Bind(new(repository.SignalRepository), new(*repository.SignalRepositoryImpl)),
)

// ServiceProvider Service 层依赖
var ServiceProvider = wire.NewSet(
	service.NewSignalService,
	wire.Bind(new(service.SignalService), new(*service.SignalServiceImpl)),
)

// APIProvider API 层依赖
var APIProvider = wire.NewSet(
	api.NewSignalAPI,
)

// ProviderSet 整合所有依赖
var ProviderSet = wire.NewSet(
	DBProvider,
	ServerProvider,
	RepositoryProvider,
	ServiceProvider,
	APIProvider,
)
