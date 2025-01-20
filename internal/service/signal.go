package service

import (
	"context"
	"errors"
	"fmt"
	"go-gin-gorm-wire-viper-zap/internal/model"
	"go-gin-gorm-wire-viper-zap/internal/repository"
)

type SignalService interface {
	Create(ctx context.Context, signal *model.Signal) error
	Update(ctx context.Context, signal *model.Signal) error
	Get(ctx context.Context, id uint64) (*model.Signal, error)
	// 业务方法
	ProcessSignal(ctx context.Context, id uint64) error
	Test(ctx context.Context) error
}

type SignalServiceImpl struct {
	signalRepo repository.SignalRepository
}

func NewSignalService(signalRepo repository.SignalRepository) SignalService {
	return &SignalServiceImpl{signalRepo: signalRepo}
}

func (s *SignalServiceImpl) Create(ctx context.Context, signal *model.Signal) error {
	return s.signalRepo.Create(ctx, signal)
}

func (s *SignalServiceImpl) Update(ctx context.Context, signal *model.Signal) error {
	return s.signalRepo.Update(ctx, signal)
}

func (s *SignalServiceImpl) Get(ctx context.Context, id uint64) (*model.Signal, error) {
	return s.signalRepo.Get(ctx, id)
}

func (s *SignalServiceImpl) ProcessSignal(ctx context.Context, id uint64) error {
	signal, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	if signal.Processed {
		return errors.New("signal already processed")
	}

	signal.Processed = true

	return s.signalRepo.Update(ctx, signal)
}

func (s *SignalServiceImpl) Test(ctx context.Context) error {
	fmt.Println("job test")
	return nil
}
