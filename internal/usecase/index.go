// index.go
// 调用repo进行数据持久化

package usecase

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
)

var _ (domain.IndexUsecase) = (*IndexUsecase)(nil)

type IndexUsecase struct {
	repo domain.IndexRepo
}

func NewIndexUsecase(repo domain.IndexRepo) *IndexUsecase {
	return &IndexUsecase{repo: repo}
}

func (uc *IndexUsecase) List(ctx context.Context, page, pageSize int) ([]*domain.Index, error) {
	// TODO: redis

	// db
	return uc.repo.ListIndex(ctx, page, pageSize)
}

func (uc *IndexUsecase) Create(ctx context.Context, index *domain.Index) error {
	return uc.repo.CreateIndex(ctx, index)
}

func (uc *IndexUsecase) Update(ctx context.Context, id int32, index *domain.Index) (int32, error) {
	return uc.repo.UpdateIndex(ctx, id, index)
}

func (uc *IndexUsecase) Delete(ctx context.Context, id int32) error {
	return uc.repo.DeleteIndex(ctx, id)
}
