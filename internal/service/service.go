package service

import "github.com/samb233/easyblog/internal/domain"

type BlogService struct {
	indexUsecase   domain.IndexUsecase
	contentUsecase domain.ContentUsecase
}

func NewBlogService(indexUsecase domain.IndexUsecase, contentUsecase domain.ContentUsecase) *BlogService {
	return &BlogService{
		indexUsecase:   indexUsecase,
		contentUsecase: contentUsecase,
	}
}
