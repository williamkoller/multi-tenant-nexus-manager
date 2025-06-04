package domain

import (
	"time"
)

type BaseEntity struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *BaseEntity) Initialize() {
	if b.ID == "" {
		b.ID = GetGenerator().Generate()
	}
	now := time.Now()
	if b.CreatedAt.IsZero() {
		b.CreatedAt = now
	}
	b.UpdatedAt = now
}

func (b *BaseEntity) GetID() string {
	b.Initialize()
	return b.ID
}
