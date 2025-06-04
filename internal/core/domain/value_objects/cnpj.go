package value_objects

import (
	"fmt"
	"regexp"
)

// CNPJ - Value Object para CNPJ brasileiro
type CNPJ struct {
	value string
}

func NewCNPJ(cnpj string) (CNPJ, error) {
	cnpj = regexp.MustCompile(`\D`).ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return CNPJ{}, fmt.Errorf("CNPJ must have 14 digits")
	}

	if !isValidCNPJ(cnpj) {
		return CNPJ{}, fmt.Errorf("invalid CNPJ: %s", cnpj)
	}

	return CNPJ{value: cnpj}, nil
}

func (c CNPJ) String() string {
	return c.value
}

func (c CNPJ) Formatted() string {
	if len(c.value) != 14 {
		return c.value
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s",
		c.value[:2], c.value[2:5], c.value[5:8], c.value[8:12], c.value[12:])
}

func isValidCNPJ(cnpj string) bool {
	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i := 0; i < 12; i++ {
		sum += int(cnpj[i]-'0') * weights1[i]
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}

	if int(cnpj[12]-'0') != digit1 {
		return false
	}

	sum = 0
	for i := 0; i < 13; i++ {
		sum += int(cnpj[i]-'0') * weights2[i]
	}
	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}

	return int(cnpj[13]-'0') == digit2
}
