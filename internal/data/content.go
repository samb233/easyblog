// content.go
// 实现ContentRepo接口

package data

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
)

type ContentRepo struct {
	data *Data
}

func NewContentRepo(data *Data) *ContentRepo {
	return &ContentRepo{
		data: data,
	}
}

var _ domain.ContentRepo = (*ContentRepo)(nil)

func (cr *ContentRepo) GetContent(ctx context.Context, id int32) (*domain.Content, error) {
	p, err := cr.data.db.Content.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Content{
		ID:        p.ID,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (cr *ContentRepo) CreateContent(ctx context.Context, content *domain.Content) (int32, error) {
	p, err := cr.data.db.Content.Create().
		SetContent(content.Content).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return p.ID, nil
}

func (cr *ContentRepo) UpdateContent(ctx context.Context, id int32, content *domain.Content) error {
	p, err := cr.data.db.Content.Get(ctx, id)
	if err != nil {
		return err
	}

	_, err = p.Update().
		SetContent(content.Content).
		Save(ctx)

	return err
}
