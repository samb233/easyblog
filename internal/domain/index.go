// index.go
// 文章目录index的领域模型定义

package domain

import (
	"context"
	"time"
)

type Index struct {
	ID        int32
	ContentID int32
	Title     string
	Desc      string
	View      int32
	Attr      int32
	CreatedAt time.Time
	UpdatedAt time.Time
	State     int8
}

// Index Repository
// Index 持久化层接口定义
type IndexRepo interface {
	// db
	ListIndex(ctx context.Context) ([]*Index, error)
	CreateIndex(ctx context.Context, index *Index) error
	UpdateIndex(ctx context.Context, id int32, index *Index) error
	DeleteIndex(ctx context.Context, id int32) error

	// TODO: redis
}

// Index Usecase
// Index 业务逻辑层接口定义
type IndexUsecase interface {
	List(ctx context.Context) ([]*Index, error)
	Create(ctx context.Context, index *Index) error
	Update(ctx context.Context, id int32, index *Index) error
	Delete(ctx context.Context, id int32) error
}
