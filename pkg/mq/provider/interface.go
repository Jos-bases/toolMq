package provider

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"toolMq/pkg/mq/conf"
)

type MqFace interface {
	CreateExchange(exchange *conf.Exchange) error
	DeleteExchange(exchange *conf.Exchange) error
	CreateQueue(queue *conf.Queue) error
	DeleteQueue(queue *conf.Queue) error
	QueueBind(key string, exchange *conf.Exchange, queue *conf.Queue) error
	QueueUnBind(key string, exchange *conf.Exchange, queue *conf.Queue) error
	Subscribe(key string, queue *conf.Queue, callback func(<-chan amqp.Delivery, string))
	SendMessage(key string, exchange *conf.Exchange, date []byte) error
	Close() error
	ChannelClose() error
}
