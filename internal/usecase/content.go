// content.go
// Content 业务逻辑层实现

package usecase

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
)

type ContentUsecase struct {
	repo domain.ContentRepo
}

func NewContentUsecase(repo domain.ContentRepo) *ContentUsecase {
	return &ContentUsecase{
		repo: repo,
	}
}

func (uc *ContentUsecase) Get(ctx context.Context, id int32) (*domain.Content, error) {
	// redis

	// db
	return uc.repo.GetContent(ctx, id)
}

func (uc *ContentUsecase) Create(ctx context.Context, content *domain.Content) (int32, error) {
	return uc.repo.CreateContent(ctx, content)
}

func (uc *ContentUsecase) Update(ctx context.Context, id int32, content *domain.Content) error {
	return uc.repo.UpdateContent(ctx, id, content)
}
