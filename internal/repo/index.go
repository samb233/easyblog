// index.go
// 实现domain中定义的IndexRepo接口

package repo

import (
	"context"
	"time"

	"github.com/samb233/easyblog/internal/domain"
	"github.com/samb233/easyblog/internal/repo/ent"
	"github.com/samb233/easyblog/internal/repo/ent/index"
	"github.com/samb233/easyblog/pkg/log"
)

type IndexRepo struct {
	repo *Repo
	log  log.Logger
}

func NewIndexRepo(repo *Repo, logger log.Logger) *IndexRepo {
	log := log.WithFields(logger, log.Fields{
		"file": "repo/index.go",
	})
	return &IndexRepo{
		repo: repo,
		log:  log,
	}
}

var _ domain.IndexRepo = (*IndexRepo)(nil)

func (ir *IndexRepo) ListIndex(ctx context.Context, page, pageSize int) ([]*domain.Index, error) {
	ps, err := ir.repo.db.Index.
		Query().
		Limit(page).
		Offset(pageSize * (page - 1)).
		Where(index.State(1)).
		Order(ent.Desc("id")).
		All(ctx)
	if err != nil {
		return nil, err
	}

	rv := make([]*domain.Index, 0)
	for _, p := range ps {
		rv = append(rv, &domain.Index{
			ID:        p.ID,
			ContentID: p.ContentID,
			Title:     p.Title,
			Desc:      p.Desc,
			View:      p.View,
			Attr:      p.Attr,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return rv, nil

}

// TODO: 事务
func (ir *IndexRepo) CreateIndex(ctx context.Context, index *domain.Index) error {
	_, err := ir.repo.db.Index.Create().
		SetContentID(index.ContentID).
		SetTitle(index.Title).
		SetDesc(index.Desc).
		SetAttr(index.Attr).
		Save(ctx)
	return err
}

func (ir *IndexRepo) UpdateIndex(ctx context.Context, id int32, index *domain.Index) error {
	p, err := ir.repo.db.Index.Get(ctx, id)
	if err != nil {
		return err
	}

	_, err = p.Update().
		SetTitle(index.Title).
		SetDesc(index.Desc).
		SetAttr(index.Attr).
		SetView(index.View).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return err
}

func (ir *IndexRepo) DeleteIndex(ctx context.Context, id int32) error {
	p, err := ir.repo.db.Index.Get(ctx, id)
	if err != nil {
		return err
	}

	_, err = p.Update().
		SetState(0).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	return err
}
