package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: 写完业务逻辑
// ListBlog的页码功能

func (h *handler) ListBlog(c *gin.Context) {
	// 获取页码
	pageQuery := c.DefaultQuery("pages", "1")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}

	indexes, err := h.blogService.ListBlog(c.Request.Context(), page)
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
