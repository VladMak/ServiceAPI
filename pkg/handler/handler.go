package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/VladMak/ServiceAPI/pkg/service"
	"github.com/VladMak/ServiceAPI/internal/chat"
)

/*
Структура обработчиков
Тут происходит инициализация роутов, по которым будет работать HTTP сервер. Создаем API для фронтэнд приложения. Подойдет как для сайта, так и для мобильного приложения сразу же.
*/
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	chat := chat.Chat()

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

			items := api.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:id", h.getItemById)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}

		chat := api.Group("/chat")
		{
			chat.POST("/", chat.Test)
			chat.GET("/getMessage")
			chat.PUT("/sendMessage")
		}
	}

	return router
}
