package httpServer

import (
	"25ugleyBot/internal/controllers/botController"
	"25ugleyBot/internal/delivery/http/routers/botRouter"
	"25ugleyBot/internal/services/botService"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start(
	host, token string,
	chatId int64,
	port int,
) error {
	// Create server
	server := gin.Default()

	// Create reposirories
	// Create services
	bs, errBotService := botService.New(token, chatId)
	if errBotService != nil {
		return errBotService
	}
	// Create controllers
	bc := botController.New(bs)
	// Add routers
	botRouter.New(server, bc)

	// Start server
	return server.Run(
		fmt.Sprintf("%s:%d", host, port),
	)
}
