package service

import "context"

// Blog用于传递消息
type Blog struct {
	Title     string
	Desc      string
	Content   string
	Attr      int32
	CreatedAt int64
	UpdatedAt int64
}

// 获取博客列表
func (bs *BlogService) ListBlog(ctx context.Context, page int) ([]*Blog, error) {
	ps, err := bs.indexUsecase.List(ctx, page, bs.pageSize)
	if err != nil {
		bs.log.Errorf(err.Error())
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
