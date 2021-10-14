package service

import (
	"context"

	"github.com/samb233/easyblog/internal/domain"
)

// Blog用于传递消息
type Blog struct {
	ID        int32  `form:"id"`
	Title     string `form:"title" valid:"Required; MaxSize(50)"`
	Desc      string `form:"desc" valid:"MaxSize(255)"`
	Content   string `form:"content" valid:"MaxSize(65535)"`
	Attr      int32  `form:"attr"`
	CreatedAt int64
	UpdatedAt int64
}

// 获取博客列表
func (bs *BlogService) ListBlog(ctx context.Context, page int) ([]*Blog, error) {
	ps, err := bs.indexUsecase.List(ctx, page, bs.pageSize)
	if err != nil {
		bs.log.Error(err.Error())
		return nil, err
	}

	blogs := make([]*Blog, 0)
	for _, p := range ps {
		blogs = append(blogs, &Blog{
			Title:     p.Title,
			Desc:      p.Desc,
			Attr:      p.Attr,
			CreatedAt: p.CreatedAt.Unix(),
			UpdatedAt: p.UpdatedAt.Unix(),
		})
	}
	return blogs, nil
}

// 获取某一篇博客
func (bs *BlogService) GetBlog(ctx context.Context, id int) (*Blog, error) {
	content, err := bs.contentUsecase.Get(ctx, int32(id))
	if err != nil {
		bs.log.Error(err.Error())
		return nil, err
	}

	return &Blog{
		Content: content.Content,
	}, nil
}

// 创建博客
func (bs *BlogService) CreateBlog(ctx context.Context, blog *Blog) error {
	content := &domain.Content{
		Content: blog.Content,
	}

	contentID, err := bs.contentUsecase.Create(ctx, content)
	if err != nil {
		// TODO: handler error
		bs.log.Error(err.Error())
		return err
	}

	index := &domain.Index{
		ContentID: contentID,
		Title:     blog.Title,
		Desc:      blog.Desc,
		Attr:      blog.Attr,
	}
	if err := bs.indexUsecase.Create(ctx, index); err != nil {
		bs.log.Error(err.Error())
		return err
	}

	return nil
}

// 更新博客
func (bs *BlogService) UpdateBlog(ctx context.Context, id int, blog *Blog) error {
	index := &domain.Index{
		Title: blog.Title,
		Desc:  blog.Desc,
		Attr:  blog.Attr,
	}
	contentID, err := bs.indexUsecase.Update(ctx, int32(id), index)
	if err != nil {
		bs.log.Error(err.Error())
		return err
	}

	content := &domain.Content{
		Content: blog.Content,
	}

	err = bs.contentUsecase.Update(ctx, contentID, content)
	if err != nil {
		bs.log.Error(err.Error())
		return err
	}

	return nil
}
