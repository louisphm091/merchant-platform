package handler

import (
	"merchant-platform/merchant-service/internal/application/command"
	"merchant-platform/merchant-service/internal/application/query"
	"merchant-platform/merchant-service/internal/controller/dto/base"
	"merchant-platform/merchant-service/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	loginHandler           *command.AdminLoginHandler
	listMerchantsHandler   *query.ListMerchantsHandler
	approveMerchantHandler *command.ApproveMerchantHandler
}

func NewAdminHandler(
	loginHandler *command.AdminLoginHandler,
	listMerchantsHandler *query.ListMerchantsHandler,
	approveMerchantHandler *command.ApproveMerchantHandler,
) *AdminHandler {
	return &AdminHandler{
		loginHandler:           loginHandler,
		listMerchantsHandler:   listMerchantsHandler,
		approveMerchantHandler: approveMerchantHandler,
	}
}

func (h *AdminHandler) Login(c *gin.Context) {

	var req command.AdminLoginCommand

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body",
			response.NewError("INVALID_REQUEST_BODY", err.Error()))
		return
	}

	result, err := h.loginHandler.Handle(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "failed to login",
			response.NewError("INVALID_CREDENTIALS", err.Error()))
		return
	}

	response.Success(c, http.StatusOK, "admin login successfully", result)
}

func (h *AdminHandler) ListMerchants(c *gin.Context) {
	result, err := h.listMerchantsHandler.Handle(c.Request.Context())

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch merchants",
			response.NewError("FETCH_MERCHANTS_FAILED", err.Error()))
		return
	}

	response.Success(c, http.StatusOK, "merchants fetched successfully", result)
}

func (h *AdminHandler) ApproveMerchant(c *gin.Context) {
	merchantID := c.Param("id")

	result, err := h.approveMerchantHandler.Handle(c.Request.Context(), base.BaseRequest[command.ApproveMerchantCommand]{
		Data: command.ApproveMerchantCommand{
			MerchantID: merchantID,
		},
	})
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to approve merchant",
			response.NewError("APPROVE_MERCHANT_FAILED", err.Error()))
		return
	}

	response.Success(c, http.StatusOK, "merchant approved successfully", result)
}
