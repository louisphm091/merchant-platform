package entity

import (
	"errors"
	"merchant-platform/merchant-service/internal/domain/auditing"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	id           uuid.UUID
	email        string
	passwordHash string
	fullName     string
	isActive     bool
	auditing     auditing.Auditing
}

func NewAdmin(email string, passwordHash string, fullName string) (*Admin, error) {

	email = strings.TrimSpace(strings.ToLower(email))
	fullName = strings.TrimSpace(fullName)

	if email == "" {
		return nil, errors.New("email is required")
	}

	if passwordHash == "" {
		return nil, errors.New("password hash is required")
	}

	if fullName == "" {
		return nil, errors.New("full name is required")
	}

	now := time.Now()

	return &Admin{
		id:           uuid.New(),
		email:        email,
		passwordHash: passwordHash,
		fullName:     fullName,
		isActive:     true,
		auditing: auditing.Auditing{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}, nil
}

func Rehydrate(id uuid.UUID, email string, passwordHash string, fullName string, isActive bool, auditing auditing.Auditing) *Admin {
	return &Admin{
		id:           id,
		email:        email,
		passwordHash: passwordHash,
		fullName:     fullName,
		isActive:     isActive,
		auditing:     auditing,
	}
}

func (a *Admin) ID() uuid.UUID {
	return a.id
}

func (a *Admin) Email() string {
	return a.email
}

func (a *Admin) PasswordHash() string {
	return a.passwordHash
}

func (a *Admin) FullName() string {
	return a.fullName
}

func (a *Admin) IsActive() bool {
	return a.isActive
}

func (a *Admin) Auditing() auditing.Auditing {
	return a.auditing
}
