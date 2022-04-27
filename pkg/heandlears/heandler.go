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
func (h Handler) GetRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{

		posts := api.Group("/post")
		{
			posts.GET("/:item_id", h.GetPostById)
			posts.GET("/list/:page/:limit", h.GetList)
			posts.POST("", h.Create)
			posts.PUT("/:item_id", h.Update)
			posts.DELETE("/:item_id", h.Delete)
		}
	}
	return router
}
