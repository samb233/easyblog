package server

import (
	"net/http"

	api "github.com/samb233/easyblog/api/v1"
	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/service"
)

func NewHTTPServer(conf *conf.Server, service *service.BlogService) *http.Server {
	srv := &http.Server{
		Addr: conf.Addr,
	}

	api.RegisterGinServer(srv, service)
	return srv
}
