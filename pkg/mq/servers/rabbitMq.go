package servers

import (
	"context"
	"fmt"
	"github.com/Jos-bases/toolMq/pkg/mq/conf"
	"github.com/Jos-bases/toolMq/pkg/mq/provider"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
)

type RabbitMq struct {
	conn *amqp.Connection
	cann *amqp.Channel
}


func (t *RabbitMq) NewConn(conf *conf.MqConf) *RabbitMq {
	new(sync.Once).Do(func() {
		conn, err := amqp.Dial(conf.NewDsn())
		if err != nil {
			panic(any(fmt.Sprintf("amqp.Dial [ERROR] : %v", err)))
		}
		t.conn = conn
		t.cann, err = conn.Channel()
		if err != nil {
			panic(any(fmt.Sprintf("amqp.conn.Channel [ERROR] : %v", err)))
		}
	})

	return t
}

func (t *RabbitMq) CreateExchange(exchange *conf.Exchange) error {

	return t.cann.ExchangeDeclare(exchange.Name, string(exchange.Types), exchange.Durable, exchange.AutoDelete, exchange.Internal, exchange.NoWait, exchange.Args)
}

func (t *RabbitMq) DeleteExchange(exchange *conf.Exchange) error {

	return t.cann.ExchangeDelete(exchange.Name, exchange.IfUnused, exchange.NoWait)
}

func (t *RabbitMq) CreateQueue(queue *conf.Queue) error {

	_, err := t.cann.QueueDeclare(queue.Name, queue.Durable, queue.AutoDelete, queue.Internal, queue.NoWait, queue.Args)

	return err
}

func (t *RabbitMq) DeleteQueue(queue *conf.Queue) error {

	_, err := t.cann.QueueDelete(queue.Name, queue.IfUnused, queue.IfEmpty, queue.NoWait)

	return err
}

func (t *RabbitMq) QueueBind(key string, exchange *conf.Exchange, queue *conf.Queue) error {

	return t.cann.QueueBind(queue.Name, key, exchange.Name, queue.NoWait, queue.Args)
}

func (t *RabbitMq) QueueUnBind(key string, exchange *conf.Exchange, queue *conf.Queue) error {

	return t.cann.QueueUnbind(queue.Name, key, exchange.Name, queue.Args)
}

func (t *RabbitMq) Subscribe(subscribe *conf.Subscribe, queue *conf.Queue, callback func(<-chan provider.MqDelivery)) {

	response, err := t.cann.Consume(queue.Name, subscribe.Name, subscribe.AutoAck, subscribe.Exclusive, subscribe.NoLocal, subscribe.NoWait, subscribe.Args)
	if err != nil {
		panic(any(err))
	}

	mqResponse := make(chan provider.MqDelivery)
	go func() {
		for delivery := range response {
			mqResponse <- provider.MqDelivery{
				ContentType:     delivery.ContentType,
				ContentEncoding: delivery.ContentEncoding,
				DeliveryMode:    delivery.DeliveryMode,
				Priority:        delivery.Priority,
				CorrelationId:   delivery.CorrelationId,
				ReplyTo:         delivery.ReplyTo,
				Expiration:      delivery.Expiration,
				MessageId:       delivery.MessageId,
				Timestamp:       delivery.Timestamp,
				Type:            delivery.Type,
				UserId:          delivery.UserId,
				AppId:           delivery.AppId,
				ConsumerTag:     delivery.ConsumerTag,
				MessageCount:    delivery.MessageCount,
				DeliveryTag:     delivery.DeliveryTag,
				Redelivered:     delivery.Redelivered,
				Exchange:        delivery.Exchange,
				RoutingKey:      delivery.RoutingKey,
				Body:            delivery.Body,
			}
		}
	}()

	callback(mqResponse)
}

func (t *RabbitMq) SendMessage(key string, exchange *conf.Exchange, data []byte) error {

	_, err := t.cann.PublishWithDeferredConfirmWithContext(context.Background(), exchange.Name, key, false, false, amqp.Publishing{
		Body: data,
	})

	return err
}

func (t *RabbitMq) Ack(tag uint64, multiple bool) error {
	return t.cann.Ack(tag, multiple)
}

func (t *RabbitMq) Nack(tag uint64, multiple bool, requeue bool) error {
	return t.cann.Nack(tag, multiple, requeue)
}

func (t *RabbitMq) Close() error {
	return t.conn.Close()
}

func (t *RabbitMq) ChannelClose() error {
	return t.cann.Close()
}
