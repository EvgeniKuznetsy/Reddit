package heandlears

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type httpError struct {
	Message string `json:"message"`
}

func sendHttpError(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}
func sendServerError(c *gin.Context) {
	c.AbortWithStatus(http.StatusInternalServerError)
}
