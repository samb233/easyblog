package v1

import (
	"github.com/gin-gonic/gin"
)

// TODO: 写完业务逻辑

func (h *handler) ListBlog(c *gin.Context) {
	indexes, err := h.blogService.ListIndex(c.Request.Context())
	if err != nil {
		// do some thing
		c.JSON(500, gin.H{
			"error": "error",
		})
	}

	c.JSON(200, indexes)
}
