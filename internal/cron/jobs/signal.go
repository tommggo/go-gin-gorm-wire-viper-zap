package jobs

import (
	"context"
	"go-gin-gorm-wire-viper-zap/internal/service"
	"go-gin-gorm-wire-viper-zap/pkg/cron"
)

const (
	SignalJobName = "signal_test"
)

type SignalJob struct {
	signalService service.SignalService
}

func NewSignalJob(signalService service.SignalService) cron.Job {
	return &SignalJob{
		signalService: signalService,
	}
}

func (j *SignalJob) Run(ctx context.Context) error {
	j.signalService.Test(ctx)
	return nil
}
