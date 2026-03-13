package utils

import (
	"merchant-platform/merchant-service/internal/controller/dto/base"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/constant"

	"github.com/gin-gonic/gin"
)

func WriteSuccess[T any](
	c *gin.Context,
	httpStatus int,
	requestId string,
	requestDateTime string,
	channel string,
	description string,
	data T) {
	c.JSON(httpStatus, base.BaseResponse[T]{
		RequestId:       requestId,
		RequestDateTime: requestDateTime,
		Channel:         channel,
		Result: base.Result{
			ResponseCode: constant.SUCCESS_CODE,
			Description:  description,
		},
		Data: data,
	})
}
