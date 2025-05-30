package value_objects

import (
	"fmt"
	"regexp"
)

type CPF struct {
	value string
}

func NewCPF(cpf string) (CPF, error) {
	cpf = regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return CPF{}, fmt.Errorf("CPF must have 11 digits")
	}

	if !isValidCPF(cpf) {
		return CPF{}, fmt.Errorf("invalid CPF: %s", cpf)
	}

	return CPF{value: cpf}, nil
}

func (c CPF) String() string {
	return c.value
}

func (c CPF) Formatted() string {
	if len(c.value) != 11 {
		return c.value
	}
	return fmt.Sprintf("%s.%s.%s-%s",
		c.value[:3], c.value[3:6], c.value[6:9], c.value[9:])
}

func isValidCPF(cpf string) bool {
	if cpf == "00000000000" || cpf == "11111111111" ||
		cpf == "22222222222" || cpf == "33333333333" ||
		cpf == "44444444444" || cpf == "55555555555" ||
		cpf == "66666666666" || cpf == "77777777777" ||
		cpf == "88888888888" || cpf == "99999999999" {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	digit1 := 11 - (sum % 11)
	if digit1 >= 10 {
		digit1 = 0
	}
	if int(cpf[9]-'0') != digit1 {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	digit2 := 11 - (sum % 11)
	if digit2 >= 10 {
		digit2 = 0
	}
	return int(cpf[10]-'0') == digit2
}
