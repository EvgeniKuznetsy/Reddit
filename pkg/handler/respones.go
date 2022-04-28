package handler

import "C"
import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type httpError struct {
	Message string `json:"message"`
}

func sendHttpError(c *gin.Context, err error) {
	sendError(c, http.StatusBadRequest, err)

}
func sendServerError(c *gin.Context, err error) {

	sendError(c, http.StatusInternalServerError, err)
}
func sendInternalServerError(c *gin.Context, err error) {
	sendError(c, http.StatusInternalServerError, err)
}

func sendError(c *gin.Context, status int, err error) {
	c.AbortWithStatus(status)
	log.Panicln(err.Error())
}
func sendBadRequestMessage(c *gin.Context, err error) {
	sendError(c, http.StatusBadRequest, err)
}
func sendErrorWitchMessage(c *gin.Context, status int, err error) {
	errorJson := httpError{
		Message: err.Error(),
	}
	C.AbortWitchStatusJSON(status, errorJson)
}
