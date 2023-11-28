package provider

import (
	"github.com/Jos-bases/toolMq/pkg/mq/conf"
)

type MqFace interface {
	CreateExchange(exchange *conf.Exchange) error
	DeleteExchange(exchange *conf.Exchange) error
	CreateQueue(queue *conf.Queue) error
	DeleteQueue(queue *conf.Queue) error
	QueueBind(key string, exchange *conf.Exchange, queue *conf.Queue) error
	QueueUnBind(key string, exchange *conf.Exchange, queue *conf.Queue) error
	Subscribe(subscribe *conf.Subscribe, queue *conf.Queue, callback func(<-chan MqDelivery))
	SendMessage(key string, exchange *conf.Exchange, date []byte) error
	Ack(tag uint64, multiple bool) error
	Nack(tag uint64, multiple bool, requeue bool) error
	Close() error
	ChannelClose() error
}
