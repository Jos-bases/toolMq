package servers

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"toolMq/pkg/mq/conf"
)

type RabbitMq struct {
	conn *amqp.Connection
	cann *amqp.Channel
}

/*
NewConn
*/
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

/*
CreateExchange
*/
func (t *RabbitMq) CreateExchange(exchange *conf.Exchange) error {

	return t.cann.ExchangeDeclare(exchange.Name, string(exchange.Types), exchange.Durable, exchange.AutoDelete, exchange.Internal, exchange.NoWait, nil)
}

func (t *RabbitMq) DeleteExchange(exchange *conf.Exchange) error {

	return t.cann.ExchangeDelete(exchange.Name, false, false)
}

func (t *RabbitMq) CreateQueue(queue *conf.Queue) error {

	_, err := t.cann.QueueDeclare(queue.Name, queue.Durable, queue.AutoDelete, queue.Internal, queue.NoWait, nil)

	return err
}

func (t *RabbitMq) DeleteQueue(queue *conf.Queue) error {

	_, err := t.cann.QueueDelete(queue.Name, false, false, false)

	return err
}

func (t *RabbitMq) QueueBind(key string, exchange *conf.Exchange, queue *conf.Queue) error {

	return t.cann.QueueBind(queue.Name, key, exchange.Name, false, nil)
}

func (t *RabbitMq) QueueUnBind(key string, exchange *conf.Exchange, queue *conf.Queue) error {

	return t.cann.QueueUnbind(queue.Name, key, exchange.Name, nil)
}

func (t *RabbitMq) Subscribe(key string, queue *conf.Queue, callback func(<-chan amqp.Delivery, string)) {

	response, err := t.cann.Consume(queue.Name, key, true, false, false, false, nil)
	if err != nil {
		panic(any(err))
	}

	callback(response, key)
}

func (t *RabbitMq) SendMessage(key string, exchange *conf.Exchange, data []byte) error {

	_, err := t.cann.PublishWithDeferredConfirmWithContext(context.Background(), exchange.Name, key, false, false, amqp.Publishing{
		Body: data,
	})

	return err
}

func (t *RabbitMq) Close() error {

	return t.conn.Close()
}

func (t *RabbitMq) ChannelClose() error {

	return t.cann.Close()
}
