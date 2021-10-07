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
			"error": err,
		})
	}

	c.JSON(200, indexes)
}

func (h *handler) GetBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "this is just a test",
	})
}

func (h *handler) CreateBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "this is just a test",
	})
}

func (h *handler) UpdateBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "this is just a test",
	})
}

func (h *handler) DeleteBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "this is just a test",
	})
}
