package service

import (
	"errors"
	"strings"

	"github.com/louisphm091/merchant-platform/internal/config"
	"github.com/louisphm091/merchant-platform/internal/model"
	"github.com/louisphm091/merchant-platform/internal/repository"
	"github.com/louisphm091/merchant-platform/internal/utils"
)

type AdminService struct {
	adminRepo *repository.AdminRepository
	cfg       *config.Config
}

func NewAdminService(adminRepo *repository.AdminRepository, cfg *config.Config) *AdminService {
	return &AdminService{
		adminRepo: adminRepo,
		cfg:       cfg,
	}
}

type AdminLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	AccessToken string       `json:"access_token"`
	Admin       *model.Admin `json:"admin"`
}

func (s *AdminService) Login(input AdminLoginInput) (*AdminLoginResponse, error) {
	email := strings.TrimSpace(strings.ToLower(input.Email))

	admin, err := s.adminRepo.FindByEmail(email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !admin.IsActive {
		return nil, errors.New("admin account is inactive")
	}

	if !utils.CheckPasswordHash(input.Password, admin.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateAdminJWT(admin.ID.String(), admin.Email, s.cfg.JWTSecret)

	if err != nil {
		return nil, err
	}

	return &AdminLoginResponse{
		AccessToken: token,
		Admin:       admin,
	}, nil
}
