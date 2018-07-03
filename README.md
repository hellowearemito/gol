# Gol

Gol is a logger what send the given type of message to redis. We can define the targets. After the sending an other service can decode/encode and send the given messages to the correct platform.

If the package could not connect to redis, then it tries to forward the message to the fallback service. The fallback service is not a require config field.

## Getting Started

### Installing

```
dep ensure -add github.com/hellowearemito/gol
```

## Usage

```go
logService := &Service{
  Host: "localhost",
  Port: "5678",
  Path: "/logger/manager"
}
config := Config{
  ListName: "log_messages",
  Redis: Service{
    Host: "localhost",
    Port: "1234"
  },
  LogService: &logService,
}

logger, err := NewLogger(config)
if err != nil {
  panic(err)
}

message := Message{
  Type: Communication,
  Targets: []Target{Dashbot},
  Data: TestData{
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
* Elastic
* Chatbase

## Built With

* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Struct validator
* [redigo](https://github.com/gomodule/redigo) - Redis client


## Authors

* **Bence Patyi** - *Mito* - [bpatyi](https://github.com/bpatyi)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
