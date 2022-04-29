package handler

import (
	"Reddit/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input models.InputSingIn

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestMessage(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		sendErrorWitchMessage(c, err)
		return
	}

	output, err := h.services.Auth.SignIn(&input)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) SignUp(c *gin.Context) {
	var input models.InputSinUp

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestMessage(c, err)
		return
	}

	if err := input.IsValid(); err != nil {
		sendBadRequestMessage(c, err)
		return
	}

	output, err := h.service.Auth.SigUp(&input)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}
