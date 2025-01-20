package cron

import (
	"fmt"
	"go-gin-gorm-wire-viper-zap/internal/config"
	"go-gin-gorm-wire-viper-zap/internal/cron/jobs"
	"go-gin-gorm-wire-viper-zap/internal/service"
	"go-gin-gorm-wire-viper-zap/pkg/cron"
)

type CronManager struct {
	signalService service.SignalService
	config        *config.Config
}

func NewCronManager(signalService service.SignalService, config *config.Config) cron.Registrar {
	return &CronManager{
		signalService: signalService,
		config:        config,
	}
}

// Register 实现 pkg/cron.Registrar 接口
func (m *CronManager) Register(c *cron.Cron) error {
	if err := m.registerSignalJobs(c); err != nil {
		return err
	}
	return nil
}

func (m *CronManager) registerSignalJobs(c *cron.Cron) error {
	signalJob := jobs.NewSignalJob(m.signalService)
	spec, ok := m.config.Cron.Specs[jobs.SignalJobName]
	if !ok {
		return fmt.Errorf("cron spec not found for job: %s", jobs.SignalJobName)
	}

	if err := c.AddJob(jobs.SignalJobName, spec, signalJob); err != nil {
		return fmt.Errorf("register account sync job failed: %w", err)
	}
	return nil
}
