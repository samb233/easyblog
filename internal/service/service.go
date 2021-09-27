package service

import (
	"github.com/samb233/easyblog/internal/domain"
)

type BlogService struct {
	indexUsecase   domain.IndexUsecase
	contentUsecase domain.ContentUsecase
}

func NewBlogService(indexUsecase domain.IndexUsecase,
	contentUsecase domain.ContentUsecase) *BlogService {
	return &BlogService{
		indexUsecase:   indexUsecase,
		contentUsecase: contentUsecase,
	}
}

// BlogDTO
type Blog struct {
	Title   string
	Desc    string
	Content string
	Attr    int32
}

// // DTO -> DO
// func blogToIndex(blog *Blog) *domain.Index {
// 	// TODO: 判断逻辑
// 	return &domain.Index{
// 		Title: blog.Title,
// 		Desc:  blog.Desc,
// 		Attr:  blog.Attr,
// 	}
// }

// func blogToContent(blog *Blog) *domain.Content {
// 	// TODO: 判断逻辑
// 	return &domain.Content{
// 		Content: blog.Content,
// 	}
// }
