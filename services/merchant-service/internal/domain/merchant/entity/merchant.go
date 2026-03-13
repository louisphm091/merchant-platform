package entity

import (
	"errors"
	"fmt"
	"merchant-platform/merchant-service/internal/domain/auditing"
	"merchant-platform/merchant-service/internal/domain/merchant/event"
	"merchant-platform/merchant-service/internal/domain/merchant/valueobject"
	"time"

	"github.com/google/uuid"
)

type MerchantStatus string

const (
	PENDING   MerchantStatus = "PENDING"
	APPROVED  MerchantStatus = "APPROVED"
	REJECTED  MerchantStatus = "REJECTED"
	SUSPENDED MerchantStatus = "SUSPENDED"
)

type DomainEvent interface {
	EventType() string
}

type Merchant struct {
	id           uuid.UUID
	merchantCode string
	businessName string
	email        valueobject.Email
	phone        string
	webhookURL   string
	status       MerchantStatus
	auditing     auditing.Auditing
	domainEvents []DomainEvent
}

func NewMerchant(
	businessName string,
	email valueobject.Email,
	phone string,
	webhookURL string) (*Merchant, error) {

	if businessName == "" {
		return nil, errors.New("business name is required")
	}

	now := time.Now()
	m := &Merchant{
		id:           uuid.New(),
		merchantCode: fmt.Sprintf("MRC-%d", now.UnixNano()),
		businessName: businessName,
		email:        email,
		phone:        phone,
		webhookURL:   webhookURL,
		status:       PENDING,
		auditing: auditing.Auditing{
			CreatedAt: now,
		},
		domainEvents: make([]DomainEvent, 0),
	}

	m.addDomainEvent(event.MerchantRegistered{
		MerchantID:   m.id.String(),
		MerchantCode: m.merchantCode,
		BusinessName: m.BusinessName(),
		Email:        m.email.Value(),
		OccurredAt:   now,
	})

	return m, nil
}

func Rehydrate(
	id uuid.UUID,
	merchantCode string,
	businessName string,
	email valueobject.Email,
	phone string,
	webhookURL string,
	status MerchantStatus,
	auditing auditing.Auditing,
) *Merchant {
	return &Merchant{
		id:           id,
		merchantCode: merchantCode,
		businessName: businessName,
		email:        email,
		phone:        phone,
		webhookURL:   webhookURL,
		status:       status,
		auditing:     auditing,
		domainEvents: make([]DomainEvent, 0),
	}
}

func (m *Merchant) Approve() error {
	if m.status == APPROVED {
		return errors.New("merchant is already approved")
	}

	if m.status == SUSPENDED {
		return errors.New("suspended merchant can not be approved directly")
	}

	m.status = APPROVED
	m.auditing.UpdatedAt = time.Now()

	m.addDomainEvent(event.MerchantApproved{
		MerchantID:   m.id.String(),
		MerchantCode: m.merchantCode,
		OccurredAt:   m.auditing.UpdatedAt,
	})

	return nil
}

func (m *Merchant) addDomainEvent(ev DomainEvent) {
	m.domainEvents = append(m.domainEvents, ev)
}

func (m *Merchant) PullDomainEvents() []DomainEvent {
	events := m.domainEvents
	m.domainEvents = make([]DomainEvent, 0)
	return events
}

func (m *Merchant) ID() uuid.UUID {
	return m.id
}

func (m *Merchant) MerchantCode() string {
	return m.merchantCode
}

func (m *Merchant) BusinessName() string {
	return m.businessName
}

func (m *Merchant) Email() string {
	return m.email.Value()
}

func (m *Merchant) Phone() string {
	return m.phone
}

func (m *Merchant) WebhookURL() string {
	return m.webhookURL
}

func (m *Merchant) Status() MerchantStatus {
	return m.status
}

func (m *Merchant) Auditing() auditing.Auditing {
	return m.auditing
}
