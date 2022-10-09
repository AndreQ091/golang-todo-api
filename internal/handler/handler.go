package handler

import (
	"github.com/AndreQ091/golang-todo/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.GetLists)
			lists.GET("/:id", h.GetOneList)
			lists.PATCH("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)

			items := lists.Group("/:id/items/")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.GetItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", h.GetOneItem)
			items.PATCH("/:id", h.UpdateItem)
			items.DELETE("/:id", h.DeleteItem)
		}
	}
	return router
}
