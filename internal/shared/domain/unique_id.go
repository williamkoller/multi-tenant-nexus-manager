package domain

import (
	"sync"

	"github.com/google/uuid"
)

type IDGenerator interface {
	Generate() string
}

type UUIDGenerator struct{}

func (g *UUIDGenerator) Generate() string {
	return uuid.New().String()
}

type UniqueID struct {
	ID string `json:"id"`
}

func (u *UniqueID) String() string {
	return u.ID
}

var (
	generator IDGenerator
	uniqueID  *UniqueID
	once      sync.Once
)

func GetGenerator() IDGenerator {
	once.Do(func() {
		generator = &UUIDGenerator{}
		uniqueID = &UniqueID{
			ID: generator.Generate(),
		}
	})
	return generator
}

func GetUniqueID() *UniqueID {
	GetGenerator()
	return uniqueID
}

func NewUniqueID() *UniqueID {
	return GetUniqueID()
}
