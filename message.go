package gol

import "github.com/go-ozzo/ozzo-validation"

// Types and targets
const (
	System        Type = "system"
	Communication Type = "communication"
	Audit         Type = "audit"

	Dashbot  Target = "dashbot"
	Chatbase Target = "chatbase"
	Elastic  Target = "elastic"
	Sentry   Target = "sentry"
	File     Target = "file"

	Incoming Source = "incoming"
	Outgoing Source = "outgoing"
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

	// Sources contains the sources of message.
	Sources = []interface{}{
		Incoming,
		Outgoing,
	}
)

// Type represents the message's type.
type Type string

// Target represents the message's target.
type Target string

// Source represents the message's source (incoming & outgoing)
type Source string

// Message represents the log's message structure.
type Message struct {
	Type       Type
	Source     *Source
	Targets    []Target
	Data       interface{}
	Intent     *Intent
	NotHandled bool
	SessionID  *string
	Version    *string
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
			&m.Source,
			validation.In(Sources...),
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

// Intent represents the intent structure.
type Intent struct {
	Name   string  `json:"name"`
	Inputs []Input `json:"inputs"`
}

// Validate validates the intent struct.
func (i Intent) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.Inputs),
	)
}

// Input represents the input of intent.
type Input struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Validate validates the input struct.
func (i Input) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.Value, validation.Required),
	)
}
