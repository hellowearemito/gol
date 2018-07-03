package gol

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

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
	config Config
}

// NewLogger returns a new logger struct.
func NewLogger(config Config) (Logger, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	return &logger{
		config: config,
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

	connection, err := redis.Dial(tcp, l.config.Redis.Domain())
	if err == nil {
		return connection.Send(rPush, l.config.ListName, data)
	}

	if l.config.LogService == nil {
		return errors.New("redis connection does not work and LogService config is not specified")
	}

	uri := url.URL{
		Scheme: "https",
		Host:   l.config.LogService.Domain(),
		Path:   l.config.LogService.Path,
	}
	body := bytes.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, uri.String(), body)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("the logger could not communicate with log service")
	}

	return nil
}
