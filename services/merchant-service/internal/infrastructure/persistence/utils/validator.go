package utils

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

var (
	UUID_REGEX      = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	DATE_TIME_REGEX = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}[-+]\d{2}:\d{2}$`)
)

func ValidateRequestId(requestId string) error {
	if strings.TrimSpace(requestId) == "" {
		return errors.New("requestId is required")
	}

	if !UUID_REGEX.MatchString(requestId) {
		return errors.New("requestId must be in format of uuid")
	}

	return nil
}

func ValidateRequestDateTime(requestDateTime string) error {
	if strings.TrimSpace(requestDateTime) == "" {
		return errors.New("requestDateTime is required")
	}

	if !DATE_TIME_REGEX.MatchString(requestDateTime) {
		return errors.New("requestDateTime must be in format of OffsetDateTime")
	}

	_, err := time.Parse("2006-01-02T15:04:05.000-07:00", requestDateTime)
	if err != nil {
		return errors.New("requestDateTime must be in format of OffsetDateTime")
	}

	return nil

}
