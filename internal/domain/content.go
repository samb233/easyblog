// content.go
// 文章内容content的领域模型定义

package domain

import (
	"context"
	"time"
)

type Content struct {
	ID        int32
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Content Repository
// Content 持久化层接口定义
type ContentRepo interface {
	// db
	GetContent(ctx context.Context, id int32) (*Content, error)
	CreateContent(ctx context.Context, content *Content) (int32, error)
	UpdateContent(ctx context.Context, id int32, content *Content) error

	// redis
}

// Content Usecase
// Content 业务逻辑层接口定义
type ContentUsecase interface {
	Get(ctx context.Context, id int32) (*Content, error)
	Create(ctx context.Context, content *Content) (int32, error)
	Update(ctx context.Context, id int32, content *Content) error
}
