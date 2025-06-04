package value_objects

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Color - Value Object para cores hexadecimais
type Color struct {
	value string
}

func NewColor(color string) (Color, error) {
	color = strings.ToUpper(strings.TrimSpace(color))

	// Adiciona # se não tiver
	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}

	// Valida formato hexadecimal
	if !regexp.MustCompile(`^#[0-9A-F]{6}$`).MatchString(color) {
		return Color{}, fmt.Errorf("invalid color format: must be #RRGGBB")
	}

	return Color{value: color}, nil
}

func (c Color) String() string {
	return c.value
}

func (c Color) RGB() (int, int, int) {
	r, _ := strconv.ParseInt(c.value[1:3], 16, 64)
	g, _ := strconv.ParseInt(c.value[3:5], 16, 64)
	b, _ := strconv.ParseInt(c.value[5:7], 16, 64)
	return int(r), int(g), int(b)
}

func (c Color) IsLight() bool {
	r, g, b := c.RGB()
	// Fórmula para determinar se a cor é clara
	brightness := (r*299 + g*587 + b*114) / 1000
	return brightness > 128
}
