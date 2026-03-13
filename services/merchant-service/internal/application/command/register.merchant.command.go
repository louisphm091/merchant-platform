package command

import (
	"context"
	"encoding/json"
	"errors"
	"merchant-platform/merchant-service/internal/controller/dto/base"
	"merchant-platform/merchant-service/internal/domain/merchant/entity"
	"merchant-platform/merchant-service/internal/domain/merchant/repository"
	"merchant-platform/merchant-service/internal/domain/merchant/valueobject"
	"strings"
)

type EventPublisher interface {
	Publish(ctx context.Context, topic string, key string, payload []byte) error
}

type RegisterMerchantCommand struct {
	BusinessName string `json:"business_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	WebhookURL   string `json:"webhook_url"`
}

type RegisterMerchantResult struct {
	ID           string `json:"id"`
	MerchantCode string `json:"merchant_code"`
	BusinessName string `json:"business_name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Status       string `json:"status"`
}

type RegisterMerchantHandler struct {
	repo      repository.MerchantRepository
	publisher EventPublisher
}

type MerchantRegisteredEvent struct {
	EventType  string `json:"event_type"`
	MerchantID string `json:"merchant_id"`
}

func NewRegisterMerchantHandler(
	repo repository.MerchantRepository,
	publisher EventPublisher,
) *RegisterMerchantHandler {
	return &RegisterMerchantHandler{
		repo:      repo,
		publisher: publisher,
	}
}

func (h *RegisterMerchantHandler) Handle(ctx context.Context, cmd base.BaseRequest[RegisterMerchantCommand]) (*RegisterMerchantResult, error) {
	email, err := valueobject.NewEmail(cmd.Data.Email)

	if err != nil {
		return nil, err
	}

	exists, err := h.repo.ExistsByEmail(ctx, email.Value())
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("merchant with this email already exists")
	}

	merchant, err := entity.NewMerchant(
		strings.TrimSpace(cmd.Data.BusinessName),
		email,
		strings.TrimSpace(cmd.Data.Phone),
		strings.TrimSpace(cmd.Data.WebhookURL),
	)

	if err != nil {
		return nil, err
	}

	if err := h.repo.Save(ctx, merchant); err != nil {
		return nil, err
	}

	events := merchant.PullDomainEvents()

	for _, evt := range events {

		eventPayload := MerchantRegisteredEvent{
			EventType:  "merchant.registered",
			MerchantID: merchant.ID().String(),
		}

		payload, err := json.Marshal(eventPayload)

		if err != nil {
			return nil, err
		}

		_ = h.publisher.Publish(ctx, "merchant.events", merchant.ID().String(), payload)
		_ = evt
	}

	return &RegisterMerchantResult{
		ID:           merchant.ID().String(),
		MerchantCode: merchant.MerchantCode(),
		BusinessName: merchant.BusinessName(),
		Phone:        merchant.Phone(),
		Email:        merchant.Email(),
		Status:       string(merchant.Status()),
	}, nil
}
