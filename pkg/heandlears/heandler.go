package heandlears

import (
	"Reddit/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}
func GetRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{

		posts := api.Group("/post")
		{
			posts.GET("/:item_id")
		}
	}
	return router
}
