package base

import "github.com/gin-gonic/gin"

func GetRequestMeta(c *gin.Context) (string, string, string) {
	requestID, _ := c.Get("requestId")
	requestDateTime, _ := c.Get("requestDateTime")
	channel, _ := c.Get("channel")

	requestIDStr, _ := requestID.(string)
	requestDateTimeStr, _ := requestDateTime.(string)
	channelStr, _ := channel.(string)

	return requestIDStr, requestDateTimeStr, channelStr
}
