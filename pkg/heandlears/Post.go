package heandlears

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetPostById(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		//to do send code 400
		sendHttpError(c)
		return
	}
	post, err := h.service.Post.GetById(id)
	if err != nil {
		//to do send code 500
		sendServerError(c)
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetList(c *gin.Context) {
	page := c.Param("page")
	limit := c.Param("limit")
	if page == "" || limit == "" {
		sendHttpError(c)
	}

	intPage, err := strconv.Atoi(page)
	if err != nil {
		sendHttpError(c)
	}
	intlimit, err := strconv.Atoi(page)
	if err != nil {
		sendHttpError(c)
	}
	h.service.Post.GetList(page, limit)
}
