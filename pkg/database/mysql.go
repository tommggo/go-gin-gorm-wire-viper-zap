package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"go-gin-gorm-wire-viper-zap/internal/config"
	"go-gin-gorm-wire-viper-zap/pkg/logger"
)

// DB 接口定义数据库操作
type DB interface {
	GetDB() *gorm.DB
	Close() error
}

// GormDB 实现 DB 接口
type GormDB struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

// NewDB 创建数据库连接
func NewDB(cfg *config.Config) (DB, error) {
	db, err := initDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("init database failed: %v", err)
	}
	return db, nil
}

// initDB 内部初始化函数
func initDB(cfg *config.Config) (*GormDB, error) {
	// 1. 配置 GORM
	gormConfig := &gorm.Config{
		Logger: gormlogger.New(
			logger.StandardLogger(),
			gormlogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  gormlogger.Info,
				IgnoreRecordNotFoundError: true,
			},
		),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	// 2. 连接数据库
	gormDB, err := gorm.Open(mysql.Open(cfg.Database.DSN), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to database failed: %w", err)
	}

	// 3. 获取底层 *sql.DB
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("get sql.DB failed: %w", err)
	}

	// 4. 配置连接池
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)

	// 5. 测试连接
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database failed: %w", err)
	}

	logger.Info("database connected",
		logger.Int("max_open_conns", cfg.Database.MaxOpenConns),
		logger.Int("max_idle_conns", cfg.Database.MaxIdleConns),
		logger.Duration("conn_max_lifetime", cfg.Database.ConnMaxLifetime),
	)

	return &GormDB{
		db:    gormDB,
		sqlDB: sqlDB,
	}, nil
}

// GetDB 获取 gorm.DB 实例
func (g *GormDB) GetDB() *gorm.DB {
	return g.db
}

// Close 关闭数据库连接
func (g *GormDB) Close() error {
	logger.Info("closing database connection...")
	return g.sqlDB.Close()
}
