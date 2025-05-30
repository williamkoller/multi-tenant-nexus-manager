package value_objects

import "fmt"

// Percentage - Value Object para porcentagens
type Percentage struct {
	value float64 // Valor entre 0 e 100
}

func NewPercentage(value float64) (Percentage, error) {
	if value < 0 || value > 100 {
		return Percentage{}, fmt.Errorf("percentage must be between 0 and 100")
	}
	return Percentage{value: value}, nil
}

func (p Percentage) Value() float64 {
	return p.value
}

func (p Percentage) Decimal() float64 {
	return p.value / 100
}

func (p Percentage) String() string {
	return fmt.Sprintf("%.2f%%", p.value)
}

func (p Percentage) ApplyTo(amount Money) Money {
	return amount.Multiply(p.Decimal())
}

func (p Percentage) Add(other Percentage) (Percentage, error) {
	newValue := p.value + other.value
	return NewPercentage(newValue)
}

func (p Percentage) Subtract(other Percentage) (Percentage, error) {
	newValue := p.value - other.value
	return NewPercentage(newValue)
}
