package valueobject

import (
	"errors"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/constant"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(constant.EmailRegex)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	normalized := strings.TrimSpace(strings.ToLower(value))

	if normalized == "" {
		return Email{}, errors.New("email is required")
	}

	if !emailRegex.MatchString(normalized) {
		return Email{}, errors.New("invalid email foramt")
	}

	return Email{value: normalized}, nil
}

func (e Email) Value() string {
	return e.value
}
