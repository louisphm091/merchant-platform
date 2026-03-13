package dto

type RegisterMerchantRequest struct {
	BusinessName string `json:business_name`
	Email        string `json:email`
	Phone        string `json:phone`
	WebhookURL   string `json:webhook_url`
}
