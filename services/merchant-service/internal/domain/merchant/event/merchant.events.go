package event

import "time"

type MerchantRegistered struct {
	MerchantID   string
	MerchantCode string
	BusinessName string
	Email        string
	OccurredAt   time.Time
}

type MerchantApproved struct {
	MerchantID   string
	MerchantCode string
	OccurredAt   time.Time
}

func (e MerchantApproved) EventType() string {
	return "merchant.approved"
}

func (e MerchantRegistered) EventType() string {
	return "merchant.regsitered"
}
