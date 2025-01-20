package cron

import (
	"context"
	"fmt"
	"go-gin-gorm-wire-viper-zap/pkg/cache"
	"go-gin-gorm-wire-viper-zap/pkg/logger"
	"time"

	"github.com/robfig/cron/v3"
)

// Job 定时任务接口
type Job interface {
	Run(ctx context.Context) error
}

type Registrar interface {
	Register(c *Cron) error
}

// Cron 定时任务管理器
type Cron struct {
	cron  *cron.Cron
	cache cache.Cache
	jobs  map[string]cron.EntryID
	specs map[string]string
}

func New(cache cache.Cache, registrar Registrar) *Cron {
	c := &Cron{
		cron:  cron.New(cron.WithSeconds()),
		cache: cache,
		jobs:  make(map[string]cron.EntryID),
		specs: make(map[string]string),
	}

	// 注册任务
	if err := registrar.Register(c); err != nil {
		panic(fmt.Sprintf("register cron jobs failed: %v", err))
	}

	return c
}

func (c *Cron) AddJob(name string, spec string, job Job) error {
	wrappedJob := func() {
		ctx := context.Background()

		// 获取锁
		lockKey := fmt.Sprintf("lock:cron:%s", name)
		ok, err := c.cache.SetNX(ctx, lockKey, []byte("1"), time.Minute)
		if err != nil {
			logger.Error("acquire lock failed",
				logger.String("job", name),
				logger.Err(err),
			)
			return
		}
		if !ok {
			return
		}
		defer c.cache.Del(ctx, lockKey)

		// 执行任务并记录日志
		startTime := time.Now()
		err = job.Run(ctx)
		duration := time.Since(startTime)

		if err != nil {
			logger.Error("run job failed",
				logger.String("job", name),
				logger.Duration("duration", duration),
				logger.Err(err),
			)
			return
		}

		logger.Info("job completed",
			logger.String("job", name),
			logger.Duration("duration", duration),
		)
	}

	id, err := c.cron.AddFunc(spec, wrappedJob)
	if err != nil {
		return fmt.Errorf("add job failed: %w", err)
	}
	c.jobs[name] = id
	c.specs[name] = spec
	return nil
}

func (c *Cron) Start() {
	c.cron.Start()

	// 打印所有已注册的任务信息
	for name, entryID := range c.jobs {
		entry := c.cron.Entry(entryID)
		logger.Info("cron job registered",
			logger.String("job", name),
			logger.String("spec", c.specs[name]),
			logger.Time("next", entry.Next),
		)
	}

	logger.Info("cron scheduler started",
		logger.Int("jobs_count", len(c.jobs)),
	)
}

func (c *Cron) Stop() {
	c.cron.Stop()
	logger.Info("cron stopped")
}
