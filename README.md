# Application

## Project Introduction

```go
@appname   toolMq
@detailed  Provide a portable tool for operating Mq
@frame     RabbitMq...
```

## Application Catalog

```go
├── README.md
├── go.mod
├── go.sum
└── pkg
└── mq
    ├── conf       
    ├── provider   
    └── servers    
```

## Using help

```go
package main

import (
	"fmt"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	"toolMq/pkg/mq/conf"
	"toolMq/pkg/mq/provider"
	"toolMq/pkg/mq/servers"
)

func main() {

	// Initialize connection
	server := provider.MqFace(
		new(servers.RabbitMq).NewConn(&conf.MqConf{
			Username: "You Username",
			Password: "You Password",
			Ip:       "You Server Ip Address",
			Port:     "You port",
		}),
	)

	// Create Exchange
	exchange := &conf.Exchange{
		Name:       "text_exchange",
		Types:      conf.Topic,
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
	}
	if err := server.CreateExchange(exchange); err != nil {
		log.Println("server.CreateExchange [ERROR] : ", err)
		return
	}

	// Create Queue
	queue := &conf.Queue{
		Name:       "test_queue",
		Types:      conf.Classic,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
	}
	if err := server.CreateQueue(queue); err != nil {
		log.Println("server.CreateQueue [ERROR] : ", err)
		return
	}

	// Bind Queue
	key := "test_top"
	if err := server.QueueBind(key, exchange, queue); err != nil {
		log.Println("server.QueueBind [ERROR] : ", err)
		return
	}

	// Subscribe Queue
	server.Subscribe(key, queue, func(deliveries <-chan amqp.Delivery, s string) {
		for delivery := range deliveries {
			fmt.Sprintf("%v Bind %v Received the news：%v\n", key, queue.Name, string(delivery.Body))
		}
	})
}
```