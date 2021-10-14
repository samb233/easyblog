package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/samb233/easyblog/internal/service"
)

// TODO: 写完业务逻辑

// 获取博客清单
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
		ReplyJson(c, http.StatusInternalServerError, -1, err.Error(), nil)
	}

	ReplyJson(c, http.StatusOK, 0, "", indexes)
}

// 获取博客内容
func (h *handler) GetBlog(c *gin.Context) {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil || blogID < 1 {
		ReplyJson(c, http.StatusBadRequest, -1, "bad request", nil)
	}

	blog, err := h.blogService.GetBlog(c.Request.Context(), blogID)
	if err != nil {
		ReplyJson(c, http.StatusInternalServerError, -1, err.Error(), nil)
	}

	ReplyJson(c, http.StatusOK, 0, "", blog)

}

// 新建博客
func (h *handler) CreateBlog(c *gin.Context) {
	blog := &service.Blog{}
	err := BindAndValid(c, blog)
	if err != nil {
		ReplyJson(c, http.StatusBadRequest, -1, err.Error(), nil)
	}

	err = h.blogService.CreateBlog(c.Request.Context(), blog)
	if err != nil {
		ReplyJson(c, http.StatusInternalServerError, -1, "未知错误", nil)
	}

	c.JSON(200, gin.H{
		"message": "create success",
	})
}

// 更新博客
func (h *handler) UpdateBlog(c *gin.Context) {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil || blogID < 1 {
		ReplyJson(c, http.StatusBadRequest, -1, "bad request", nil)
	}

	blog := &service.Blog{}
	err = BindAndValid(c, blog)
	if err != nil {
		ReplyJson(c, http.StatusBadRequest, -1, err.Error(), nil)
	}

	err = h.blogService.UpdateBlog(c.Request.Context(), blogID, blog)
	if err != nil {
		ReplyJson(c, http.StatusInternalServerError, -1, err.Error(), nil)
	}

	ReplyJson(c, http.StatusOK, 0, "update success", nil)

}

func (h *handler) DeleteBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "this is just a test",
	})
}

// 绑定并验证
func BindAndValid(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		return err
	}

	valid := validation.Validation{}
	ok, err := valid.Valid(form)
	if err != nil {
		return err
	}

	if !ok {
		var errStr string
		for _, err := range valid.Errors {
			errStr = fmt.Sprintf("errKey: %s, err: %s; ", err.Key, err.String())
		}
		return fmt.Errorf("valid error: %s", errStr)
	}

	return nil
}

// 返回值定义
type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回json
func ReplyJson(c *gin.Context, httpCode, code int, msg string, data interface{}) {
	c.JSON(httpCode, Reply{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
