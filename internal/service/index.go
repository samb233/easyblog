// index.go
// 定义Index服务
// DO到DTO的转换
// 数据校验、组装，不涉及复杂逻辑

package service

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
)

// IndexRequest
type IndexRequest struct {
	ContentID int32
	Title     string
	Desc      string
	Attr      int32
	CreatedAt int64
	UpdatedAt int64
}

func (bs *BlogService) CreateIndex(ctx context.Context, ireq *IndexRequest) error {
	index := &domain.Index{
		ContentID: ireq.ContentID,
		Title:     ireq.Title,
		Desc:      ireq.Desc,
		Attr:      ireq.Attr,
	}
	if err := bs.indexUsecase.Create(ctx, index); err != nil {
		// TODO: 处理error
		return err
	}

	return nil
}

// TODO: 分页
func (bs *BlogService) ListIndex(ctx context.Context) ([]*IndexRequest, error) {
	ps, err := bs.indexUsecase.List(ctx)
	if err != nil {
		return nil, err
	}

	indexes := make([]*IndexRequest, 0)
	for _, p := range ps {
		indexes = append(indexes, &IndexRequest{
			ContentID: p.ContentID,
			Title:     p.Title,
			Desc:      p.Desc,
			Attr:      p.Attr,
			CreatedAt: p.CreatedAt.Unix(),
			UpdatedAt: p.UpdatedAt.Unix(),
		})
	}

	return indexes, nil
}

// update index
func (bs *BlogService) UpdateIndex(ctx context.Context, id int32, ireq *IndexRequest) error {
	index := &domain.Index{
		Title: ireq.Title,
		Desc:  ireq.Desc,
		Attr:  ireq.Attr,
	}

	return bs.indexUsecase.Update(ctx, id, index)
}

// delete index
func (bs *BlogService) DeleteIndex(ctx context.Context, id int32) error {
	return bs.indexUsecase.Delete(ctx, id)
}
