package gol

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gomodule/redigo/redis"
)

const (
	tcp   string = "tcp"
	rPush string = "RPUSH"
)

// FallbackLogger represents a fallback solution if something went wrong, but you want to handle it somehow
type FallbackLogger interface {
	Error(args ...interface{})
}

// Logger represents the logger interface.
type Logger interface {
	Log(message Message)
}

// logger represents the logger.
type logger struct {
	config          Config
	redisPool       *redis.Pool
	fallbackLoggers []FallbackLogger
}

// NewLogger returns a new logger struct.
func NewLogger(config Config, redisPool *redis.Pool, fallbackLoggers ...FallbackLogger) (Logger, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	return &logger{
		config:          config,
		redisPool:       redisPool,
		fallbackLoggers: fallbackLoggers,
	}, nil
}

// Log sends the log message to redis.
func (l *logger) Log(message Message) {
	err := message.Validate()
	if err != nil {
		l.fallbackLog("message.Validate():", err, message)
	}

	data, err := json.Marshal(message)
	if err != nil {
		l.fallbackLog("json.Marshal():", err, message)
	}

	err = l.redisPool.Get().Send(rPush, l.config.ListName, data)
	if err == nil {
		if err = l.redisPool.Close(); err != nil {
			l.fallbackLog("l.redisPool.Close():", err)
		}
		return
	}

	l.fallbackLog("l.redisPool.Get().Send():", err, string(data))

	if l.config.LogService == nil {
		l.fallbackLog("redis connection is not working and LogService config is not defined:", err, string(data))
		return
	}

	uri := url.URL{
		Scheme: "https",
		Host:   l.config.LogService.Domain(),
		Path:   l.config.LogService.Path,
	}
	body := bytes.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, uri.String(), body)
	if err != nil {
		l.fallbackLog("http.NewRequest():", err, string(data))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		l.fallbackLog("http.DefaultClient.Do():", err, string(data))
		return
	}

	if resp.StatusCode != http.StatusOK {
		l.fallbackLog("the logger could not communicate with log service:", err, string(data))
	}
}

func (l *logger) fallbackLog(args ...interface{}) {
	for _, lg := range l.fallbackLoggers {
		lg.Error(args...)
	}
}
