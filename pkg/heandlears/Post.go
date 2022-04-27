package heandlears

import "C"
import (
	"Reddit/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetPostById(c *gin.Context) {

	id := c.Param("item_id")
	if id == "" {
		//to do send code 400
		sendHttpError(c, errors.New("Invalid item_id"))
		return
	}
	post, err := h.service.Post.GetById(id)
	if err != nil {
		//to do send code 500
		sendServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetList(c *gin.Context) {
	page := c.Param("page")
	limit := c.Param("limit")
	if page == "" || limit == "" {
		sendHttpError(c, errors.New("Invalid item_id"))
	}

	intPage, err := strconv.Atoi(page)
	if err != nil {
		sendHttpError(c, errors.New("Invalid item_id"))
	}
	intLimit, err := strconv.Atoi(page)
	if err != nil {
		sendHttpError(c, errors.New("Invalid item_id"))
	}
	if intLimit <= 0 || intPage <= 0 {
		sendHttpError(c, errors.New("Invalid item_id"))
	}
	output, err := h.service.Post.GetList(intPage, intLimit)
	if err != nil {
		sendHttpError(c, errors.New("Invalid item_id"))
		return
	}
	c.JSON(http.StatusOK, output)
}

func (h *Handler) Create(c *gin.Context) {
	var input models.InputPost
	if err := c.BindJSON(&input); err != nil {
		sendHttpError(c, errors.New("Invalid item_id"))
		return
	}
	output, err := h.service.Post.Create(&input)
	if err != nil {
		sendServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, output)

}
func (h *Handler) Update(c *gin.Context) {
	var input models.InputUpdatesPost
	id := c.Param("item_id")
	if id == "" {
		sendHttpError(c, errors.New("Invalid item_id"))
		return
	}
	input.Id = id
	if err := c.BindJSON(&input); err != nil {
		sendHttpError(c, errors.New("Invalid item_id"))
		return
	}
	err := h.service.Post.Update(&input)
	if err != nil {
		sendServerError(c, err)
		return
	}
	c.Status(http.StatusOK)
	return
}
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		sendHttpError(c, errors.New("Invalid item_id"))
		return
	}

	if err := h.service.Post.Delete(id); err != nil {
		sendServerError(c, err)
		return
	}
	c.Status(http.StatusOK)
	return
}
