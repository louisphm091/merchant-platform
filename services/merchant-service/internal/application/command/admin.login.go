package command

import (
	"context"
	"errors"
	"merchant-platform/merchant-service/internal/domain/admin/repository"
	"merchant-platform/merchant-service/internal/infrastructure/config"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/security/jwt"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/utils"
	"strings"
)

type AdminLoginCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	AccessToken string      `json:"access_token"`
	Admin       interface{} `json:"admin"`
}

type AdminLoginHandler struct {
	adminRepository repository.AdminRepository
	cfg             *config.Config
}

func NewAdminLoginHandler(
	adminRepository repository.AdminRepository,
	cfg *config.Config) *AdminLoginHandler {
	return &AdminLoginHandler{
		adminRepository: adminRepository,
		cfg:             cfg,
	}
}

func (h *AdminLoginHandler) Handle(ctx context.Context, cmd AdminLoginCommand) (*AdminLoginResponse, error) {
	admin, err := h.adminRepository.FindByEmail(ctx, strings.TrimSpace(strings.ToLower(cmd.Email)))

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !admin.IsActive() {
		return nil, errors.New("admin account is inactive")
	}

	if !utils.CheckPasswordHash(cmd.Password, admin.PasswordHash()) {
		return nil, errors.New("invalid password")
	}

	token, err := jwt.GenerateAdminJWT(admin.ID().String(), admin.Email(), h.cfg.JWTSecret)

	if err != nil {
		return nil, err
	}

	return &AdminLoginResponse{
		AccessToken: token,
		Admin: map[string]interface{}{
			"id":       admin.ID().String(),
			"email":    admin.Email(),
			"fullName": admin.FullName(),
			"isActive": admin.IsActive(),
			"auditing": admin.Auditing(),
		},
	}, nil

}
