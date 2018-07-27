# Gol

Gol is a logger what send the given message to redis. You can define message targets targets. After the sending an other service can decode/encode and send the given messages to the correct platform.

If the package can't connect to redis, it tries to send the message to the fallback service via HTTP post request. Fallback service is not a required.

## Getting Started

### Installing 

```
go get -u github.com/hellowearemito/gol
```
> or use dependency manager ❤️

## Usage

```go
logService := &gol.Service{
  Host: "localhost",
  Port: "5678",
  Path: "/logger/manager"
}
config := gol.Config{
  ListName: "log_messages",
  LogService: &logService,
}

pool := &redis.Pool{
  MaxIdle:     3,
  IdleTimeout: 240 * time.Second,
  Dial: func() (redis.Conn, error) {
    return redis.Dial("tcp", "localhost:6379")
  },
}

logger, err := gol.NewLogger(config, pool, someFallbacklogger, otherFallbackLogger)
if err != nil {
  panic(err)
}

message := gol.Message{
  Type: gol.Communication,
  Targets: []gol.Target{gol.Dashbot},
  Data: struct{Test string}{
    Test: "data",
  }
}

err = logger.Log(message)
if err != nil{
  panic(err)
}
```

### Types

* System
* Communication
* Audit

### Targets

* Dashbot
* Chatbase
* Elastic
* Logstash
* Sentry
* File

## Built With

* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Struct validator
* [redigo](https://github.com/gomodule/redigo) - Redis client


## Authors

* **Bence Patyi** - *Mito* - [bpatyi](https://github.com/bpatyi)
* **Attila Sumi** - *Mito* - [sumia01](https://github.com/sumia01)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
