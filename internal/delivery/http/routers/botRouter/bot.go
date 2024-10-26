package botRouter

import (
	"25ugleyBot/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func New(e *gin.Engine, controller botController) {
	router := e.Group("/api/bot")
	router.Use(middlewares.CheckBotKey)
	{
		router.POST("/booking", controller.SendBooking)
	}
}
