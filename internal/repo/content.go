// content.go
// 实现ContentRepo接口

package repo

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
	"github.com/samb233/easyblog/pkg/log"
)

type ContentRepo struct {
	repo *Repo
	log  log.Logger
}

func NewContentRepo(repo *Repo, logger log.Logger) *ContentRepo {
	log := log.WithFields(logger, log.Fields{
		"file": "repo/content.go",
	})
	return &ContentRepo{
		repo: repo,
		log:  log,
	}
}

var _ domain.ContentRepo = (*ContentRepo)(nil)

func (cr *ContentRepo) GetContent(ctx context.Context, id int32) (*domain.Content, error) {
	p, err := cr.repo.db.Content.Get(ctx, id)
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
	p, err := cr.repo.db.Content.Create().
		SetContent(content.Content).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return p.ID, nil
}

func (cr *ContentRepo) UpdateContent(ctx context.Context, id int32, content *domain.Content) error {
	p, err := cr.repo.db.Content.Get(ctx, id)
	if err != nil {
		return err
	}

	_, err = p.Update().
		SetContent(content.Content).
		Save(ctx)

	return err
}
