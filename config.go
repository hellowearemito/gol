package golw

import (
	"github.com/go-ozzo/ozzo-validation"
)

// Config represents the redis configuration for golw logger.
type Config struct {
	ListName string
	Host     string
	Port     string
}

// Validate validates the struct.
func (c Config) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.ListName, validation.Required),
		validation.Field(&c.Host, validation.Required),
		validation.Field(&c.Port, validation.Required),
	)
}
