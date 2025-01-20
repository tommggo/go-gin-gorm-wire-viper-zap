package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin-gorm-wire-viper-zap/internal/config"
	"go-gin-gorm-wire-viper-zap/pkg/cache"
	"go-gin-gorm-wire-viper-zap/pkg/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func New(cfg *config.Config) (cache.Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:            cfg.Redis.Addr,
		Password:        cfg.Redis.Password,
		DB:              cfg.Redis.DB,
		PoolSize:        cfg.Redis.MaxOpenConns,
		MinIdleConns:    cfg.Redis.MaxIdleConns,
		ConnMaxLifetime: cfg.Redis.ConnMaxLifetime,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping redis failed: %w", err)
	}

	logger.Info("redis connected",
		logger.String("addr", cfg.Redis.Addr),
		logger.Int("max_open_conns", cfg.Redis.MaxOpenConns),
		logger.Int("max_idle_conns", cfg.Redis.MaxIdleConns),
		logger.Duration("conn_max_lifetime", cfg.Redis.ConnMaxLifetime),
	)

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	return r.client.Get(ctx, key).Bytes()
}

func (r *Redis) Set(ctx context.Context, key string, value []byte) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *Redis) SetEX(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Redis) GetObject(ctx context.Context, key string, value interface{}) error {
	data, err := r.Get(ctx, key)
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, value)
}

func (r *Redis) SetObject(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Set(ctx, key, data)
}

func (r *Redis) SetObjectEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.SetEX(ctx, key, data, expiration)
}

func (r *Redis) SetNX(ctx context.Context, key string, value []byte, expiration time.Duration) (bool, error) {
	return r.client.SetNX(ctx, key, value, expiration).Result()
}

func (r *Redis) Close() error {
	logger.Info("closing redis connection...")
	return r.client.Close()
}
