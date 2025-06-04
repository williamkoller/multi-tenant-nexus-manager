package value_objects

import (
	"fmt"
	"regexp"
	"strings"
)

// Slug - Value Object para URLs amigáveis
type Slug struct {
	value string
}

func NewSlug(text string) (Slug, error) {
	slug := strings.ToLower(strings.TrimSpace(text))

	// Remove acentos (simplificado)
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "ã", "a", "â", "a", "ä", "a",
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"í", "i", "ì", "i", "î", "i", "ï", "i",
		"ó", "o", "ò", "o", "õ", "o", "ô", "o", "ö", "o",
		"ú", "u", "ù", "u", "û", "u", "ü", "u",
		"ç", "c", "ñ", "n",
	)
	slug = replacer.Replace(slug)

	// Remove caracteres especiais e substitui espaços por hífens
	slug = regexp.MustCompile(`[^a-z0-9\s-]`).ReplaceAllString(slug, "")
	slug = regexp.MustCompile(`\s+`).ReplaceAllString(slug, "-")
	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	if slug == "" {
		return Slug{}, fmt.Errorf("invalid slug: cannot be empty after processing")
	}

	if len(slug) > 100 {
		return Slug{}, fmt.Errorf("slug too long: maximum 100 characters")
	}

	return Slug{value: slug}, nil
}

func (s Slug) String() string {
	return s.value
}

func (s Slug) IsEmpty() bool {
	return s.value == ""
}
