package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisphm091/merchant-platform/internal/service"
	"github.com/louisphm091/merchant-platform/pkg/response"
)

type AdminHandler struct {
	adminService *service.AdminService
}

func NewAdminHandler(adminService *service.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (a *AdminHandler) Login(c *gin.Context) {
	var req service.AdminLoginInput

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	result, err := a.adminService.Login(req)

	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	response.Success(c, http.StatusOK, "Admin Login successfully", result)
}
