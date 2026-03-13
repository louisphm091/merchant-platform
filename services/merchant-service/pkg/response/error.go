package response

type ErrorDetail struct {
	Code    string `json:"code"`
	Details string `json:"details,omitempty"`
}

func NewError(code string, details string) ErrorDetail {
	return ErrorDetail{
		Code:    code,
		Details: details,
	}
}
