// index.go
// 定义Index服务
// DO到DTO的转换
// 数据校验、组装，不涉及复杂逻辑

package service

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
)

// Index DTO
type Index struct {
	ContentID int32
	Title     string
	Desc      string
	Attr      int32
	CreatedAt int64
	UpdatedAt int64
}

func (bs *BlogService) CreateIndex(ctx context.Context, index *Index) error {
	indexDO := &domain.Index{
		ContentID: index.ContentID,
		Title:     index.Title,
		Desc:      index.Desc,
		Attr:      index.Attr,
	}
	if err := bs.indexUsecase.Create(ctx, indexDO); err != nil {
		// TODO: 处理error
		return err
	}

	return nil
}

// TODO: 分页
func (bs *BlogService) ListIndex(ctx context.Context) ([]*Index, error) {
	ps, err := bs.indexUsecase.List(ctx)
	if err != nil {
		return nil, err
	}

	indexes := make([]*Index, 0)
	for _, p := range ps {
		indexes = append(indexes, &Index{
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
