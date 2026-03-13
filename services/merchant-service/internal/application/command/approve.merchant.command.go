package command

import (
	"context"
	"encoding/json"
	"errors"
	"merchant-platform/merchant-service/internal/controller/dto/base"
	"merchant-platform/merchant-service/internal/domain/merchant/repository"
)

type ApproveMerchantCommand struct {
	MerchantID string `json:"merchant_id" binding:"required"`
}

type ApproveMerchantResponse struct {
	ID           string `json:"id"`
	MerchantCode string `json:"merchant_code"`
	BusinessName string `json:"business_name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
}

type ApproveMerchantHandler struct {
	repo      repository.MerchantRepository
	publisher EventPublisher
}

type ApproveMerchantEvent struct {
	EventType  string `json:"event_type"`
	MerchantID string `json:"merchant_id"`
}

func NewApproveMerchantHandler(
	repo repository.MerchantRepository,
	publisher EventPublisher) *ApproveMerchantHandler {
	return &ApproveMerchantHandler{
		repo:      repo,
		publisher: publisher,
	}
}

func (a *ApproveMerchantHandler) Handle(ctx context.Context, cmd base.BaseRequest[ApproveMerchantCommand]) (*ApproveMerchantResponse, error) {

	merchant, err := a.repo.FindById(ctx, cmd.Data.MerchantID)

	if err != nil {
		return nil, errors.New("merchant not found")
	}

	if err := merchant.Approve(); err != nil {
		return nil, err
	}

	if err := a.repo.Update(ctx, merchant); err != nil {
		return nil, err
	}

	events := merchant.PullDomainEvents()

	for _, evt := range events {
		payload, _ := json.Marshal(map[string]interface{}{
			"event_type":    evt.EventType(),
			"merchant_id":   merchant.ID().String(),
			"merchant_code": merchant.MerchantCode(),
			"status":        merchant.Status(),
		})

		_ = a.publisher.Publish(ctx, "merchant.events", merchant.ID().String(), payload)
	}

	return &ApproveMerchantResponse{
		ID:           merchant.ID().String(),
		MerchantCode: merchant.MerchantCode(),
		BusinessName: merchant.BusinessName(),
		Email:        merchant.Email(),
		Status:       string(merchant.Status()),
	}, nil
}
