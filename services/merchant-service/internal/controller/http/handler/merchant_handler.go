package handler

import (
	"merchant-platform/merchant-service/internal/application/command"
	"merchant-platform/merchant-service/internal/controller/dto/base"
	"merchant-platform/merchant-service/internal/controller/dto/merchant"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	registerHandler *command.RegisterMerchantHandler
}

func NewMerchantHandler(registerHandler *command.RegisterMerchantHandler) *MerchantHandler {
	return &MerchantHandler{
		registerHandler: registerHandler,
	}
}

func (h *MerchantHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "merchant-service is running",
	})
}

func (h *MerchantHandler) Register(c *gin.Context) {
	var req base.BaseRequest[command.RegisterMerchantCommand]
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	result, err := h.registerHandler.Handle(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.Result{
			ResponseCode: constant.INVALID_REQUEST,
			Description:  "Invalid request",
		})
		return
	}

	c.JSON(http.StatusCreated, base.BaseResponse[merchant.RegisterMerchantResponse]{
		RequestId:       req.RequestId,
		RequestDateTime: req.RequestDateTime,
		Channel:         req.Channel,
		Result: base.Result{
			ResponseCode: constant.SUCCESS_CODE,
			Description:  "Registered merchant successfully",
		},
		Data: merchant.RegisterMerchantResponse{
			ID:           result.ID,
			BusinessName: result.BusinessName,
			Phone:        result.Phone,
			Email:        result.Email,
			Status:       result.Status,
		},
	})
}
