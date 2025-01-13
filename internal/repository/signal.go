package repository

import (
	"context"
	"go-gin-gorm-wire-viper-zap/internal/model"
	"go-gin-gorm-wire-viper-zap/pkg/database"
)

type SignalRepository interface {
	Create(ctx context.Context, signal *model.Signal) error
	Update(ctx context.Context, signal *model.Signal) error
	Get(ctx context.Context, id uint64) (*model.Signal, error)
	// 特殊方法
	GetUnprocessed(ctx context.Context) ([]*model.Signal, error)
}

type SignalRepositoryImpl struct {
	*BaseRepository[model.Signal]
}

func NewSignalRepository(db database.DB) *SignalRepositoryImpl {
	return &SignalRepositoryImpl{
		BaseRepository: NewBaseRepository(db, &model.Signal{}),
	}
}

// 只需要实现特殊方法
func (r *SignalRepositoryImpl) GetUnprocessed(ctx context.Context) ([]*model.Signal, error) {
	var signals []*model.Signal
	err := r.withContext(ctx).Where("processed = ?", false).Find(&signals).Error
	return signals, err
}
