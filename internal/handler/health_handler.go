package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisphm091/merchant-platform/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	response.Success(c, http.StatusOK, "Service is running", gin.H{
		"service": "merchant-platform",
		"status":  "ok",
	})
}
