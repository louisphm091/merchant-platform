package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/louisphm091/merchant-platform/internal/model"
	"github.com/louisphm091/merchant-platform/internal/repository"
)

// private final MerchantRepository merchantRepository;
type MerchantService struct {
	merchantRepo *repository.MerchantRepository
}

func NewMerchantService(merchantRepo *repository.MerchantRepository) *MerchantService {
	return &MerchantService{
		merchantRepo: merchantRepo,
	}
}

type RegisterMerchantInput struct {
	BusinessName string `json:business_name`
	Email        string `json:email`
	Phone        string `json:phone`
	WebhookURL   string `json:webhook_url`
}

func (s *MerchantService) Register(input RegisterMerchantInput) (*model.Merchant, error) {

	existingMerchant, _ := s.merchantRepo.FindByEmail(input.Email)
	if existingMerchant != nil {
		return nil, fmt.Errorf("Merchant with this email already exists")
	}

	merchant := &model.Merchant{
		ID:           uuid.New(),
		MerchantCode: generateMerchantCode(),
		BusinessName: strings.TrimSpace(input.BusinessName),
		Email:        strings.TrimSpace(strings.ToLower(input.Email)),
		Phone:        strings.TrimSpace(input.Phone),
		WebhookURL:   strings.TrimSpace(input.WebhookURL),
		Status:       model.MerchantStatusPending,
	}

	if err := s.merchantRepo.Create(merchant); err != nil {
		return nil, err
	}

	return merchant, nil
}

func (s *MerchantService) List() ([]model.Merchant, error) {
	return s.merchantRepo.FindAll()
}

func generateMerchantCode() string {
	return fmt.Sprintf("MRC-%d", time.Now().UnixNano())
}
