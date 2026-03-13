package merchant

type RegisterMerchantResponse struct {
	ID           string `json:"id"`
	BusinessName string `json:"business_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Status       string `json:"status"`
}
