package domain_user

import (
	"time"

	"github.com/williamkoller/multi-tenant-nexus-manager/internal/shared/domain"
	"github.com/williamkoller/multi-tenant-nexus-manager/internal/shared/domain/value_objects"
)

type User struct {
	domain.BaseAggregateRoot
	ID       string              `json:"id"`
	Email    value_objects.Email `json:"email"`
	CPF      value_objects.CPF   `json:"cpf"`
	Phone    value_objects.Phone `json:"phone"`
	IsActive bool                `json:"is_active"`
}

func (u *User) Activate() {
	u.IsActive = true

	event := domain.NewBaseDomainEvent(
		"user.activated",
		u.GetID(),
		map[string]interface{}{
			"email":        u.Email.String(),
			"activated_at": time.Now().UTC().Format(time.RFC3339),
		},
	)
	u.RaiseDomainEvent(event)
}

func NewUser(u *User) (*User, error) {
	return &User{
		ID:    u.GetID(),
		Email: u.Email,
		CPF:   u.CPF,
		Phone: u.Phone,
	}, nil
}
