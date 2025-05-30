package value_objects

import (
	"fmt"
	"regexp"
)

// Phone - Value Object para telefone brasileiro
type Phone struct {
	value string
}

func NewPhone(phone string) (Phone, error) {
	phone = regexp.MustCompile(`\D`).ReplaceAllString(phone, "")

	// Telefone brasileiro: 10 ou 11 dígitos
	if len(phone) < 10 || len(phone) > 11 {
		return Phone{}, fmt.Errorf("phone must have 10 or 11 digits")
	}

	// Validação básica do DDD
	if len(phone) >= 2 {
		ddd := phone[:2]
		validDDDs := map[string]bool{
			"11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true,
			"21": true, "22": true, "24": true, "27": true, "28": true,
			"31": true, "32": true, "33": true, "34": true, "35": true, "37": true, "38": true,
			"41": true, "42": true, "43": true, "44": true, "45": true, "46": true,
			"47": true, "48": true, "49": true,
			"51": true, "53": true, "54": true, "55": true,
			"61": true, "62": true, "63": true, "64": true, "65": true, "66": true, "67": true, "68": true, "69": true,
			"71": true, "73": true, "74": true, "75": true, "77": true, "79": true,
			"81": true, "82": true, "83": true, "84": true, "85": true, "86": true, "87": true, "88": true, "89": true,
			"91": true, "92": true, "93": true, "94": true, "95": true, "96": true, "97": true, "98": true, "99": true,
		}

		if !validDDDs[ddd] {
			return Phone{}, fmt.Errorf("invalid area code: %s", ddd)
		}
	}

	return Phone{value: phone}, nil
}

func (p Phone) String() string {
	return p.value
}

func (p Phone) Formatted() string {
	if len(p.value) == 10 {
		return fmt.Sprintf("(%s) %s-%s", p.value[:2], p.value[2:6], p.value[6:])
	}
	if len(p.value) == 11 {
		return fmt.Sprintf("(%s) %s-%s", p.value[:2], p.value[2:7], p.value[7:])
	}
	return p.value
}

func (p Phone) IsMobile() bool {
	// Celular tem 11 dígitos e o terceiro dígito é 9
	return len(p.value) == 11 && p.value[2] == '9'
}
