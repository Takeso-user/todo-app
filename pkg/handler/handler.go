package handler

import (
	"fmt"
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
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
	api := router.Group("/api", h.userIdentity)
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

func (h *Handler) userIdentity(context *gin.Context) {
	header := context.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(context, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 {
		newErrorResponse(context, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.service.Authorization.ParseToken(headerSplit[1])
	if err != nil {
		newErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}
	context.Set(userCtx, userId)
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
	logrus.Infof("!!!user id: %d", context.GetInt(userCtx))
	id, err := getUserId(context)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, fmt.Sprintf("user id= %d not found", id))
		return
	}
	var input todoapp.TodoList
	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	id, err = h.service.TodoList.Create(id, input)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	logrus.Infof("!!!getUserId: user id: %d", id)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("!user id= %d not found", id))
		return 0, fmt.Errorf("error: user id not found")
	}
	idInt, ok := id.(int)
	logrus.Infof("idInt=`%v`", idInt)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("!user id= %d is of invalid type", id))
		return 0, fmt.Errorf("user id is of invalid type")
	}
	return idInt, nil
}

type getAllListsResponse struct {
	Data []todoapp.TodoList `json:"data"`
}

func (h *Handler) getAllLists(context *gin.Context) {
	logrus.Infof("!!!user id: %d", context.GetInt(userCtx))
	id, err := getUserId(context)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, fmt.Sprintf("user id= %d not found", id))
		return
	}

	lists, err := h.service.TodoList.GetAll(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, getAllListsResponse{Data: lists})
}

func (h *Handler) getListById(context *gin.Context) {
	logrus.Infof("!!!user id: %d", context.GetInt(userCtx))
	userId, err := getUserId(context)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, fmt.Sprintf("user id= %d not found", userId))
		return
	}
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	list, err := h.service.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, list)
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
