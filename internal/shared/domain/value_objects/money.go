package value_objects

import (
	"fmt"
	"strings"
)

// Money - Value Object para valores monetários
type Money struct {
	amount   int64  // Valor em centavos
	currency string // Código da moeda (BRL, USD, etc.)
}

func NewMoney(amount float64, currency string) Money {
	return Money{
		amount:   int64(amount * 100), // Converte para centavos
		currency: strings.ToUpper(currency),
	}
}

func NewMoneyFromCents(cents int64, currency string) Money {
	return Money{
		amount:   cents,
		currency: strings.ToUpper(currency),
	}
}

func (m Money) Amount() float64 {
	return float64(m.amount) / 100
}

func (m Money) AmountInCents() int64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("cannot add different currencies: %s and %s", m.currency, other.currency)
	}
	return Money{amount: m.amount + other.amount, currency: m.currency}, nil
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("cannot subtract different currencies: %s and %s", m.currency, other.currency)
	}
	return Money{amount: m.amount - other.amount, currency: m.currency}, nil
}

func (m Money) Multiply(factor float64) Money {
	return Money{
		amount:   int64(float64(m.amount) * factor),
		currency: m.currency,
	}
}

func (m Money) Divide(divisor float64) Money {
	if divisor == 0 {
		return m
	}
	return Money{
		amount:   int64(float64(m.amount) / divisor),
		currency: m.currency,
	}
}

func (m Money) IsPositive() bool {
	return m.amount > 0
}

func (m Money) IsNegative() bool {
	return m.amount < 0
}

func (m Money) IsZero() bool {
	return m.amount == 0
}

func (m Money) GreaterThan(other Money) bool {
	return m.currency == other.currency && m.amount > other.amount
}

func (m Money) LessThan(other Money) bool {
	return m.currency == other.currency && m.amount < other.amount
}

func (m Money) Equal(other Money) bool {
	return m.currency == other.currency && m.amount == other.amount
}

func (m Money) String() string {
	return fmt.Sprintf("%.2f %s", m.Amount(), m.currency)
}

func (m Money) FormattedBRL() string {
	if m.currency != "BRL" {
		return m.String()
	}
	return fmt.Sprintf("R$ %.2f", m.Amount())
}