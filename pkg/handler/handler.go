package handler

import (
	"github.com/Takeso-user/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			item := lists.Group("/:id/items")
			{
				item.POST("/", h.createItem)
				item.GET("/", h.getAllItems)
				item.GET("/:item_id", h.getItemById)
				item.PUT("/:item_id", h.updateItem)
				item.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}

func (h *Handler) signUp(context *gin.Context) {

}

func (h *Handler) signIn(context *gin.Context) {

}

func (h *Handler) createList(context *gin.Context) {

}

func (h *Handler) getAllLists(context *gin.Context) {

}

func (h *Handler) getListById(context *gin.Context) {

}

func (h *Handler) deleteList(context *gin.Context) {

}

func (h *Handler) updateList(context *gin.Context) {

}

func (h *Handler) createItem(context *gin.Context) {

}

func (h *Handler) getAllItems(context *gin.Context) {

}

func (h *Handler) getItemById(context *gin.Context) {

}

func (h *Handler) updateItem(context *gin.Context) {

}

func (h *Handler) deleteItem(context *gin.Context) {

}
