package handler

import (
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(context *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	context.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}

func (h *Handler) signUp(context *gin.Context) {
	var input todoapp.User
	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(context *gin.Context) {
	var input signInInput
	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
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
