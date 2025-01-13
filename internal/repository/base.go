package repository

import (
	"context"
	"go-gin-gorm-wire-viper-zap/internal/errors"
	"go-gin-gorm-wire-viper-zap/pkg/database"

	"gorm.io/gorm"
)

// BaseRepository 提供通用的 CRUD 操作
type BaseRepository[T any] struct {
	db database.DB
}

// NewBaseRepository 创建基础仓储实例
func NewBaseRepository[T any](db database.DB, model *T) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

func (r *BaseRepository[T]) withContext(ctx context.Context) *gorm.DB {
	return r.db.GetDB().WithContext(ctx).Model(new(T))
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.withContext(ctx).Create(entity).Error
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.withContext(ctx).Updates(entity).Error
}

func (r *BaseRepository[T]) Get(ctx context.Context, id uint64) (*T, error) {
	var entity T
	if err := r.withContext(ctx).First(&entity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(errors.ErrDB, err)
	}

	return &entity, nil
}

func (r *BaseRepository[T]) List(ctx context.Context, page, size int) ([]*T, error) {
	var entities []*T
	err := r.withContext(ctx).Offset((page - 1) * size).Limit(size).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}
