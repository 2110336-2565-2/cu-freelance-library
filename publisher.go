package gosdk

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Publisher interface {
	Publish(exchangeName string, topic string, body any) error
}

func NewPublisher(conn *amqp.Connection) (Publisher, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &publisher{
		channel: channel,
	}, nil
}

type publisher struct {
	logger  *Logger
	channel *amqp.Channel
}

func (p *publisher) Publish(exchangeName string, topic string, message any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	if err != nil {
		p.logger.
			Error(err).
			Keyword("exchange_name", exchangeName).
			Keyword("topic", topic).
			Msg("error while convert message to json")
		return err
	}

	if err := p.channel.PublishWithContext(ctx,
		exchangeName,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		p.logger.
			Error(err).
			Keyword("exchange_name", exchangeName).
			Keyword("topic", topic).
			Msg("error while publish message to rabbitmq")
		return err
	}

	p.logger.
		Info().
		Keyword("exchange_name", exchangeName).
		Keyword("topic", topic).
		Msg("successfully publish message to rabbitmq")

	return nil
}
