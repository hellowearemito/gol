package golw

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const (
	tcp   string = "tcp"
	rPush string = "RPUSH"
)

// Logger represents the logger interface.
type Logger interface {
	Log(message Message) error
}

// logger represents the logger.
type logger struct {
	config     Config
	connection redis.Conn
}

// NewLogger returns a new logger struct.
func NewLogger(config Config) (Logger, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	connection, err := redis.Dial(tcp, fmt.Sprintf("%v:%v", config.Host, config.Port))
	if err != nil {
		return nil, err
	}

	return &logger{
		config:     config,
		connection: connection,
	}, nil
}

// Log sends the log message to redis.
func (l *logger) Log(message Message) error {
	err := message.Validate()
	if err != nil {
		return err
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return l.connection.Send(rPush, l.config.ListName, data)
}
