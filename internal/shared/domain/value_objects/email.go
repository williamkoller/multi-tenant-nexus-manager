package value_objects

import (
	"fmt"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return Email{}, fmt.Errorf("invalid email format: %s", email)
	}

	return Email{value: email}, nil
}

func (e Email) String() string {
	return e.value
}

func (e Email) IsEmpty() bool {
	return e.value == ""
}
