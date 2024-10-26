package botRouter

import "github.com/gin-gonic/gin"

type botController interface {
	SendBooking(ctx *gin.Context)
}
