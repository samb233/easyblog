package service

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
)

type ContentRequest struct {
	Content   string
	CreatedAt int64
	UpdatedAt int64
}

func (bs *BlogService) CreateContent(ctx context.Context, creq *ContentRequest) (int32, error) {
	content := &domain.Content{
		Content: creq.Content,
	}

	return bs.contentUsecase.Create(ctx, content)
}

func (bs *BlogService) UpdateContent(ctx context.Context, id int32, creq *ContentRequest) error {
	content := &domain.Content{
		Content: creq.Content,
	}

	return bs.contentUsecase.Update(ctx, id, content)
}

func (bs *BlogService) GetContent(ctx context.Context, id int32) (*ContentRequest, error) {
	content, err := bs.contentUsecase.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	creq := &ContentRequest{
		Content:   content.Content,
		CreatedAt: content.CreatedAt.Unix(),
		UpdatedAt: content.UpdatedAt.Unix(),
	}

	return creq, nil
}
