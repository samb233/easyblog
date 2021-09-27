package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samb233/easyblog/internal/service"
)

type handler struct {
	blogService *service.BlogService
}

func RegisterGinServer(srv *http.Server, service *service.BlogService) {
	handler := &handler{
		blogService: service,
	}
	engine := gin.New()
	engine.GET("/ping", pingHandler)

	apiv1 := engine.Group("api/v1")
	{
		apiv1.GET("/article", handler.ListBlog)
	}
	srv.Handler = engine
}

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
