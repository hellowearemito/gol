# Golr

Golr is a logger what send the given type of message to redis. We can define the targets. After the sending an other service can decode/encode and send the given messages to the correct platform.

## Getting Started

### Installing

```
dep ensure -add github.com/hellowearemito/golw
```

## Usage

```go
config = Config{
  ListName: "log_messages",
  Host: "localhost",
  Port: "1234"
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

logger.Log(message)
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
