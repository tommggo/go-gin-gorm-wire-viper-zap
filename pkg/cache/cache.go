package cache

import (
	"context"
	"time"
)

// Cache 缓存接口
type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error                                     // 永不过期
	SetEX(ctx context.Context, key string, value []byte, expiration time.Duration) error         // 带过期时间
	SetNX(ctx context.Context, key string, value []byte, expiration time.Duration) (bool, error) // 分布式锁
	Del(ctx context.Context, key string) error
	GetObject(ctx context.Context, key string, value interface{}) error
	SetObject(ctx context.Context, key string, value interface{}) error                             // 永不过期
	SetObjectEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error // 带过期时间
	Close() error
}
