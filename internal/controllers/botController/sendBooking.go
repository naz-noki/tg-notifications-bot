package botController

import (
	"25ugleyBot/pkg/logger"
	"25ugleyBot/pkg/sendRespose"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *botController) SendBooking(ctx *gin.Context) {
	op := "botController/SendBooking: "
	body := new(SendBookingRequest)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(op + errBindJSON.Error())
		sendRespose.Send(ctx, http.StatusBadRequest, "error", "Not all data has been transferred to the server.", nil)
		return
	}

	errSendBooking := bc.service.SendBooking(body.Name, body.Time, body.Date, body.Phone, body.Quantity, body.Is_confirmed)
	if errSendBooking != nil {
		logger.Log.Error(op + errSendBooking.Error())
		sendRespose.Send(ctx, http.StatusInternalServerError, "error", "An error occurred while sending the message.", nil)
		return
	}

	sendRespose.Send(ctx, http.StatusOK, "success", "OK.", nil)
}
