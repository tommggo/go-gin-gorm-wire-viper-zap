//go:build wireinject
// +build wireinject

package di

import (
	"go-gin-gorm-wire-viper-zap/internal/api"
	"go-gin-gorm-wire-viper-zap/internal/config"
	"go-gin-gorm-wire-viper-zap/internal/di/provider"
	"go-gin-gorm-wire-viper-zap/pkg/database"
	"go-gin-gorm-wire-viper-zap/pkg/http"

	"github.com/google/wire"
)

// Container 包含所有依赖
type Container struct {
	Server    *http.Server
	DB        database.DB
	SignalAPI *api.SignalAPI
}

// InitializeServer 初始化服务器和所有依赖
func InitializeServer(cfg *config.Config) (*Container, error) {
	wire.Build(
		wire.Struct(new(Container), "*"),
		provider.ProviderSet,
	)
	return nil, nil
}
