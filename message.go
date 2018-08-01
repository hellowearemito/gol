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
	Logstash Target = "logstash"
	Sentry   Target = "sentry"
	File     Target = "file"

	Incoming Source = "incoming"
	Outgoing Source = "outgoing"

	Facebook Platform = "Facebook"
	SMS      Platform = "SMS"
	Web      Platform = "Web"
	Android  Platform = "Android"
	IOS      Platform = "iOS"
	Actions  Platform = "Actions"
	Alexa    Platform = "Alexa"
	Cortana  Platform = "Cortana"
	Kik      Platform = "Kik"
	Skype    Platform = "Skype"
	Twitter  Platform = "Twitter"
	Viber    Platform = "Viber"
	Telegram Platform = "Telegram"
	Slack    Platform = "Slack"
	WhatsApp Platform = "WhatsApp"
	WeChat   Platform = "WeChat"
	Line     Platform = "Line"
	Kakao    Platform = "Kakao"
	RBM      Platform = "RBM"
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
		Logstash,
		Sentry,
		File,
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

// Platform represents the message's platform.
type Platform string

// Message represents the log's message structure.
type Message struct {
	Type        Type        `json:"type"`
	Source      *Source     `json:"source"`
	Platform    *Platform   `json:"platform"`
	Targets     []Target    `json:"targets"`
	RecipientID *string     `json:"recipient_id"`
	SenderID    *string     `json:"sender_id"`
	AccessToken *string     `json:"access_token"`
	SessionID   *string     `json:"session_id"`
	MessageID   *string     `json:"message_id"`
	SentTime    time.Time   `json:"sent_time"`
	Data        interface{} `json:"data"`
	Intent      *Intent     `json:"intent"`
	NotHandled  bool        `json:"not_handled"`
	Version     *string     `json:"version"`
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

				if t == Communication {
					if s != Incoming || s != Outgoing {
						return errors.New("the source is required for the communication type")
					}

					if m.Platform == nil {
						return errors.New("the platform is required for the communication type")
					}
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
					if m.InTarget(File) || m.InTarget(Sentry) || m.InTarget(Logstash) {
						return err
					}
				case Communication:
					if m.InTarget(Dashbot) || m.InTarget(Chatbase) || m.InTarget(Logstash) {
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
					if !m.InTarget(Logstash) {
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
