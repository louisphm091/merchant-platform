package query

import (
	"context"
	merchantRepo "merchant-platform/merchant-service/internal/domain/merchant/repository"
)

type ListMerchantsHandler struct {
	repo merchantRepo.MerchantRepository
}

func NewListMerchantsHandler(repo merchantRepo.MerchantRepository) *ListMerchantsHandler {
	return &ListMerchantsHandler{repo: repo}
}

func (h *ListMerchantsHandler) Handle(ctx context.Context) ([]map[string]interface{}, error) {
	merchants, err := h.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0, len(merchants))
	for _, merchant := range merchants {
		result = append(result, map[string]interface{}{
			"id":            merchant.ID().String(),
			"merchant_code": merchant.MerchantCode(),
			"business_name": merchant.BusinessName(),
			"email":         merchant.Email(),
			"phone":         merchant.Phone(),
			"webhook_url":   merchant.WebhookURL(),
			"status":        merchant.Status(),
			"auditing":      merchant.Auditing(),
		})
	}

	return result, nil
}
