package middlewares

import (
	"25ugleyBot/pkg/configParser"
	"25ugleyBot/pkg/sendRespose"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckBotKey(ctx *gin.Context) {
	bot_token := ctx.GetHeader("bot_token")

	if bot_token != configParser.Cfg.Bot.Key {
		sendRespose.Send(ctx, http.StatusForbidden, "error", "Invalid access key to the bot.", nil)
		ctx.Abort()
		return
	}

	ctx.Next()
}
