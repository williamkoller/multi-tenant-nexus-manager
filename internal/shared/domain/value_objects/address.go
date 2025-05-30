package value_objects

import (
	"fmt"
	"regexp"
	"strings"
)

// Address - Value Object para endereço
type Address struct {
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zip_code"`
	Country    string `json:"country"`
}

func NewAddress(street, number, complement, district, city, state, zipCode, country string) (Address, error) {
	if street == "" || city == "" || state == "" || zipCode == "" {
		return Address{}, fmt.Errorf("street, city, state and zip_code are required")
	}

	// Remove caracteres especiais do CEP
	zipCode = regexp.MustCompile(`\D`).ReplaceAllString(zipCode, "")

	// Validação básica do CEP brasileiro
	if len(zipCode) != 8 {
		return Address{}, fmt.Errorf("invalid zip code format")
	}

	return Address{
		Street:     strings.TrimSpace(street),
		Number:     strings.TrimSpace(number),
		Complement: strings.TrimSpace(complement),
		District:   strings.TrimSpace(district),
		City:       strings.TrimSpace(city),
		State:      strings.ToUpper(strings.TrimSpace(state)),
		ZipCode:    zipCode,
		Country:    strings.TrimSpace(country),
	}, nil
}

func (a Address) FullAddress() string {
	parts := []string{a.Street}

	if a.Number != "" {
		parts = append(parts, a.Number)
	}
	if a.Complement != "" {
		parts = append(parts, a.Complement)
	}
	if a.District != "" {
		parts = append(parts, a.District)
	}

	parts = append(parts, a.City, a.State)

	if a.ZipCode != "" {
		parts = append(parts, a.FormattedZipCode())
	}
	if a.Country != "" {
		parts = append(parts, a.Country)
	}

	return strings.Join(parts, ", ")
}

func (a Address) FormattedZipCode() string {
	if len(a.ZipCode) == 8 {
		return fmt.Sprintf("%s-%s", a.ZipCode[:5], a.ZipCode[5:])
	}
	return a.ZipCode
}

func (a Address) IsComplete() bool {
	return a.Street != "" && a.City != "" && a.State != "" && a.ZipCode != ""
}
