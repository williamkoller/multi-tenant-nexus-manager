package value_objects

import (
	"fmt"
	"regexp"
	"strings"
)

// Code - Value Object para códigos alfanuméricos
type Code struct {
	value string
}

func NewCode(code string, minLength, maxLength int) (Code, error) {
	code = strings.ToUpper(strings.TrimSpace(code))

	if len(code) < minLength || len(code) > maxLength {
		return Code{}, fmt.Errorf("code must be between %d and %d characters", minLength, maxLength)
	}

	// Apenas letras e números
	if !regexp.MustCompile(`^[A-Z0-9]+$`).MatchString(code) {
		return Code{}, fmt.Errorf("code must contain only letters and numbers")
	}

	return Code{value: code}, nil
}

func (c Code) String() string {
	return c.value
}

func (c Code) IsEmpty() bool {
	return c.value == ""
}
