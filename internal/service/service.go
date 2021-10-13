package service

import (
	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/domain"
	"github.com/samb233/easyblog/pkg/log"
)

type BlogService struct {
	indexUsecase   domain.IndexUsecase
	contentUsecase domain.ContentUsecase

	log log.Logger

	// 每页大小
	pageSize int
}

func NewBlogService(conf *conf.Service, logger log.Logger, indexUsecase domain.IndexUsecase, contentUsecase domain.ContentUsecase) *BlogService {
	log := log.WithFields(logger, log.Fields{
		"file": "blog.go",
	})
	return &BlogService{
		indexUsecase:   indexUsecase,
		contentUsecase: contentUsecase,

		log:      log,
		pageSize: conf.Pagesize,
	}
}
