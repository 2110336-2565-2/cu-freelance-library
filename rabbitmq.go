package gosdk

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RabbitMQ interface {
	Close()
	Listen() ([]byte, error)
	CreateQueue(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args map[string]interface{}) (amqp.Queue, error)
	CreateExchange(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args map[string]interface{}) error
	BindQueueWithExchange(queueName string, key string, exchangeName string, noWait bool, args map[string]interface{}) error
	CreateMessageChannel(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args map[string]interface{}) error
	Publish(exchangeName string, topic string, message any) error
}

func NewRabbitMQ(conn *amqp.Connection) (RabbitMQ, error) {
	logger := NewLogger("rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		logger.
			Error(err).
			Msg("Error when creating channel")
		return nil, err
	}

	return &rabbitmq{
		logger:  logger,
		conn:    conn,
		channel: ch,
	}, nil
}

type rabbitmq struct {
	logger      Logger
	conn        *amqp.Connection
	channel     *amqp.Channel
	errCh       <-chan *amqp.Error
	messageChan <-chan amqp.Delivery
}

func (r *rabbitmq) Close() {
	r.channel.Close()
}

func (r *rabbitmq) Publish(exchangeName string, key string, message any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	if err != nil {
		r.logger.
			Error(err).
			Keyword("exchange_name", exchangeName).
			Keyword("key", key).
			Msg("error while convert message to json")
		return err
	}

	if err := r.channel.PublishWithContext(ctx,
		exchangeName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		r.logger.
			Error(err).
			Keyword("exchange_name", exchangeName).
			Keyword("key", key).
			Msg("error while publish message to rabbitmq")
		return err
	}

	r.logger.
		Info().
		Keyword("exchange_name", exchangeName).
		Keyword("key", key).
		Msg("successfully publish message to rabbitmq")

	return nil
}

func (r *rabbitmq) CreateQueue(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args map[string]interface{}) (amqp.Queue, error) {
	queue, err := r.channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		r.logger.
			Error(err).
			Keyword("name", name).
			Keyword("durable", durable).
			Keyword("autoDelete", autoDelete).
			Keyword("exclusive", exclusive).
			Keyword("noWait", noWait).
			Msg("error while creating queue")
		return amqp.Queue{}, err
	}

	return queue, nil
}

func (r *rabbitmq) CreateExchange(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args map[string]interface{}) error {
	r.logger.SetName(name + "-rabbitmq")

	if err := r.channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args); err != nil {
		r.logger.
			Error(err).
			Keyword("name", name).
			Keyword("kind", kind).
			Keyword("durable", durable).
			Keyword("autoDelete", autoDelete).
			Keyword("internal", internal).
			Keyword("noWait", noWait).
			Msg("error while creating exchange")
		return err
	}

	return nil
}

func (r *rabbitmq) BindQueueWithExchange(queueName string, key string, exchangeName string, noWait bool, args map[string]interface{}) error {
	if err := r.channel.QueueBind(queueName, key, exchangeName, noWait, args); err != nil {
		r.logger.
			Error(err).
			Msg("error while queue binding")
		return err
	}

	r.logger.
		Info().
		Keyword("queue_name", queueName).
		Keyword("exchange", exchangeName).
		Keyword("key", key).
		Keyword("noWait", noWait).
		Msg("binding queue")

	return nil
}

func (r *rabbitmq) CreateMessageChannel(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args map[string]interface{}) error {
	var err error
	r.messageChan, err = r.channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	if err != nil {
		r.logger.
			Error(err).
			Keyword("queue", queue).
			Keyword("consumer", consumer).
			Keyword("autoAck", autoAck).
			Keyword("exclusive", exclusive).
			Keyword("noLocal", noLocal).
			Keyword("noWait", noWait).
			Msg("Error when consuming message")
		return err
	}

	return nil
}

func (r *rabbitmq) Listen() ([]byte, error) {
	select {
	case err := <-r.errCh:
		r.logger.
			Error(err).
			Msg("error while receive message from rabbitmq")
		return nil, err
	case msg := <-r.messageChan:
		return msg.Body, nil
	}
}
