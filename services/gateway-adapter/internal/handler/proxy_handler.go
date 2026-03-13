package handler

import (
	"io"
	"merchant-platform/gateway-adapter/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProxyHandler struct {
	cfg    *config.Config
	client *http.Client
}

func NewProxyHandler(cfg *config.Config) *ProxyHandler {
	return &ProxyHandler{
		cfg:    cfg,
		client: &http.Client{},
	}
}

func (h *ProxyHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "gateway-adapter is running",
	})
}

func (h *ProxyHandler) ProxyRegisterMerchant(c *gin.Context) {
	targetURL := h.cfg.MerchantServiceURL + "/api/merchants/register"

	req, err := http.NewRequestWithContext(c.Request.Context(), http.MethodPost, targetURL, c.Request.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to create upstream request",
			"error":   err.Error(),
		})
	}

	req.Header = c.Request.Header.Clone()
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": "failed to call merchant-service",
			"error":   err.Error(),
		})
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "failed to read upstream response",
			"error":   err.Error(),
		})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}
