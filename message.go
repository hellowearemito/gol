package gol

import (
	"github.com/go-ozzo/ozzo-validation"
)

const (
	System        Type = "system"
	Communication Type = "communication"
	Audit         Type = "audit"

	Dashbot  Target = "dashbot"
	Chatbase Target = "chatbase"
	Elastic  Target = "elastic"
)

var (
	// Types contains the types of message.
	Types = []interface{}{
		System,
		Communication,
		Audit,
	}

	// Targets contains the targets of message.
	Targets = []interface{}{
		Dashbot,
		Chatbase,
		Elastic,
	}
)

// Type represents the message's type.
type Type string

// Target represents the message's target.
type Target string

// Message represents the log's message structure.
type Message struct {
	Type    Type
	Targets []Target
	Data    interface{}
}

// Validate validates the message.
func (m Message) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(
			&m.Type,
			validation.Required,
			validation.In(Types...),
		),
		validation.Field(
			&m.Targets,
			validation.Required,
			validation.In(Targets...),
		),
		validation.Field(
			&m.Data,
			validation.Required,
		),
	)
}
