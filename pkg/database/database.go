package database

import (
	"gorm.io/gorm"
)

// DB 数据库接口
type DB interface {
	GetDB() *gorm.DB
	Close() error
}
