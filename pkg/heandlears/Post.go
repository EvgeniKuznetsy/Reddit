package heandlears

import "C"
import (
	"Reddit/models"
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
	intLimit, err := strconv.Atoi(page)
	if err != nil {
		sendHttpError(c)
	}
	if intLimit <= 0 || intPage <= 0 {
		sendHttpError(c)
	}
	output, err := h.service.Post.GetList(intPage, intLimit)
	if err != nil {
		sendHttpError(c)
		return
	}
	c.JSON(http.StatusOK, output)
}

func (h *Handler) Create(c *gin.Context) {
	var input models.InputPost
	if err := c.BindJSON(&input); err != nil {
		sendHttpError(c)
		return
	}
	output, err := h.service.Post.Create(&input)
	if err != nil {
		sendServerError(c)
		return
	}
	c.JSON(http.StatusOK, output)

}
func (h *Handler) Update() (c *gin.Context) {
	var input models.InputUpdatesPost
	id := c.Param("iteam_id")
	if id == "" {
		sendHttpError(c)
		return
	}
	if err := c.BindJSON(&input); err != nil {
		sendHttpError(c)
		return
	}
	err := h.service.Post.Update(&input)
	if err != nil {
		sendServerError(c)
		return
	}
	c.Status(http.StatusOK)
	return
}
func (h *Handler) Delete() (c *gin.Context) {
	id := c.Param(("inteam_id"))
	if id == "" {
		sendHttpError(c)
		return
	}

	if err := h.service.Post.Delete(id); err != nil {
		sendServerError(c)
		return
	}
	c.Status(http.StatusOK)
	return
}
