package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisphm091/merchant-platform/internal/service"
	"github.com/louisphm091/merchant-platform/pkg/response"
)

// private final MerchantService merchantService
type MerchantHandler struct {
	merchantService *service.MerchantService
}

func NewMerchantHandler(merchantService *service.MerchantService) *MerchantHandler {
	return &MerchantHandler{
		merchantService: merchantService,
	}
}

func (h *MerchantHandler) Approve(c *gin.Context) {

	id := c.Param("id")
	merchant, err := h.merchantService.ApproveMerchant(id)

	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to approve merchant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Merchants approved successfully", merchant)
}

func (h *MerchantHandler) Reject(c *gin.Context) {
	id := c.Param("id")
	merchant, err := h.merchantService.RejectMerchant(id)

	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to reject merchant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Merchants rejected successfully", merchant)
}

func (h *MerchantHandler) Register(c *gin.Context) {

	var req service.RegisterMerchantInput

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	merchant, err := h.merchantService.Register(req)

	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to register merchant", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Merchant successfully registered", merchant)
}

func (h *MerchantHandler) List(c *gin.Context) {
	merchants, err := h.merchantService.List()

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to list all merchants", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Merchants fetched successfully", merchants)
}
