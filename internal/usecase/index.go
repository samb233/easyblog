// index.go
// Index 业务逻辑层实现
// 调用repo进行数据持久化
// 吞掉repo层传入的error

package usecase

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
)

type IndexUsecase struct {
	repo domain.IndexRepo
}

func NewIndexUsecase(repo domain.IndexRepo) *IndexUsecase {
	return &IndexUsecase{repo: repo}
}

func (uc *IndexUsecase) List(ctx context.Context) ([]*domain.Index, error) {
	// TODO: redis

	// db
	return uc.repo.ListIndex(ctx)
}

func (uc *IndexUsecase) Create(ctx context.Context, index *domain.Index) error {
	return uc.repo.CreateIndex(ctx, index)
}

func (uc *IndexUsecase) Update(ctx context.Context, id int32, index *domain.Index) error {
	return uc.repo.UpdateIndex(ctx, id, index)
}

func (uc *IndexUsecase) Delete(ctx context.Context, id int32) error {
	return uc.repo.DeleteIndex(ctx, id)
}
