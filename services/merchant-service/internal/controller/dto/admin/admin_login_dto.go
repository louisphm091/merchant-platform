package admin

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
