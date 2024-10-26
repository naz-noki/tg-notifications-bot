package sendRespose

import "github.com/gin-gonic/gin"

func Send(
	c *gin.Context,
	code int, // HTTP status code
	status string, // success | error
	msg string,
	data interface{},
) {
	c.JSON(
		code,
		gin.H{
			"status":  status,
			"message": msg,
			"data":    data,
		},
	)
}
