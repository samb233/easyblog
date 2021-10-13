// index.go
// Define index domain

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
type IndexRepo interface {
	// db
	ListIndex(ctx context.Context, page, pageSize int) ([]*Index, error)
	CreateIndex(ctx context.Context, index *Index) error
	UpdateIndex(ctx context.Context, id int32, index *Index) error
	DeleteIndex(ctx context.Context, id int32) error

	// TODO: redis
}

// Index Usecase
type IndexUsecase interface {
	List(ctx context.Context, page, pageSize int) ([]*Index, error)
	Create(ctx context.Context, index *Index) error
	Update(ctx context.Context, id int32, index *Index) error
	Delete(ctx context.Context, id int32) error
}
