package gol

import (
	"strings"

	"github.com/go-ozzo/ozzo-validation"
)

// Config represents the redis configuration for golw logger.
type Config struct {
	ListName   string
	Redis      Service
	LogService *Service
}

// Validate validates the struct.
func (c Config) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.ListName, validation.Required),
		validation.Field(&c.Redis, validation.Required),
		validation.Field(&c.LogService),
	)
}

// Service represents the service configuration
type Service struct {
	Host string
	Port string
	Path string
}

// Validate validates the struct.
func (s Service) Validate() error {
	return validation.ValidateStruct(
		&s,
		validation.Field(&s.Host, validation.Required),
		validation.Field(&s.Port, validation.Required),
	)
}

// Domain returns the service's domain.
func (s Service) Domain() string {
	return strings.Join([]string{s.Host, s.Port}, ":")
}
