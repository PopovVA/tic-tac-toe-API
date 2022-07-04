package handler

import (
	"github.com/PopovVA/tic-tac-toe-API/pkg/service"
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

	api := router.Group("/api")
	{
		games := api.Group("/games")
		{
			games.POST("/", h.CreateGame)
			games.GET("/:id", h.GetGame)
			games.DELETE("/:id", h.DeleteGame)
			games.PUT("/:id", h.UpdateGame)

			history := api.Group(":id/history")
			{
				history.GET("/", h.GetHistory)
				history.DELETE("/:historyId", h.DeleteHistory)
				history.PUT("/:historyId", h.UpdateHistory)
			}
		}
	}

	return router
}
