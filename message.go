package gol

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"

	"github.com/pkg/errors"
)

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
	Type        Type
	Source      *Source
	Targets     []Target
	RecipientID *string
	SenderID    *string
	AccessToken *string
	SessionID   *string
	MessageID   *string
	SentTime    time.Time
	Data        interface{}
	Intent      *Intent
	NotHandled  bool
	Version     *string
}

func (m *Message) InTarget(target Target) bool {
	if len(m.Targets) == 0 {
		return false
	}

	for _, t := range m.Targets {
		if t == target {
			return true
		}
	}
	return false
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
			validation.By(func(value interface{}) error {
				s, _ := value.(Source)
				t := m.Type

				if t == Communication && (s != Incoming || s != Outgoing) {
					return errors.New("the source is required for the communication type")
				}

				return nil
			}),
		),
		validation.Field(
			&m.Targets,
			validation.Required,
			validation.In(Targets...),
			validation.By(func(value interface{}) error {
				var err error
				t := m.Type

				err = errors.New("the target is not valid for: " + string(t) + " type")
				switch t {
				case System:
					if m.InTarget(File) || m.InTarget(Sentry) || m.InTarget(Elastic) {
						return err
					}
				case Communication:
					if m.InTarget(Dashbot) || m.InTarget(Chatbase) || m.InTarget(Elastic) {
						return err
					}

					if m.RecipientID == nil {
						return errors.New("the recipient id is required")
					}

					if m.SenderID == nil {
						return errors.New("the sender id is required")
					}

					if m.AccessToken == nil {
						return errors.New("the access token is required")
					}

					if m.SessionID == nil {
						return errors.New("the session id is required")
					}

					if m.MessageID == nil {
						return errors.New("the message id is required")
					}
				case Audit:
					if m.InTarget(Elastic) {
						return err
					}
				}

				return nil
			}),
		),
		validation.Field(
			&m.Data,
			validation.Required,
		),
		validation.Field(
			&m.SentTime,
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
