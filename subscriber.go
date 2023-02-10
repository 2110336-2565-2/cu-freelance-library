package gosdk

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageHandler func(ctx context.Context)

type Subscriber struct {
	ch           *amqp.Channel
	errCh        chan error
	topicList    []string
	exchangeKind string
	exchangeName string
	queueName    string
	handlerList  []MessageHandler
	logger       Logger
}

func NewSubscriber(logger Logger, connection *amqp.Connection, errCh chan error, topicList []string, exchangeKind string, exchangeName string) (*Subscriber, error) {
	ch, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	logger.SetName(exchangeName)

	return &Subscriber{
		ch:           ch,
		errCh:        errCh,
		topicList:    topicList,
		exchangeKind: exchangeKind,
		exchangeName: exchangeName,
		logger:       logger,
	}, nil
}

func (s *Subscriber) Listen() {
	msgs, err := s.initConsumer()
	if err != nil {
		s.errCh <- err
	}

	s.logger.Info().
		Keyword("queue_name", s.queueName).
		Keyword("exchange", s.exchangeName).
		Msg("Ready to accept message from RabbitMQ")

	for d := range msgs {
		s.logger.Info().
			Keyword("topic", d.RoutingKey).
			Msg("Received event")

		for _, messageHandler := range s.handlerList {
			messageHandler(context.WithValue(context.Background(), "message", d.Body))
		}
	}
}

func (s *Subscriber) Close() error {
	return s.ch.Close()
}

func (s *Subscriber) initConsumer() (<-chan amqp.Delivery, error) {
	if err := s.ch.ExchangeDeclare(
		s.exchangeName,
		s.exchangeKind,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		s.logger.Error(err).
			Keyword("exchange", s.exchangeName).
			Msg("Failed to create exchange")
		return nil, err
	}

	queue, err := s.ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		s.logger.Error(err).
			Msg("Failed to create queue")
		return nil, err
	}

	s.queueName = queue.Name

	for _, topic := range s.topicList {
		s.logger.Info().
			Keyword("queue_name", queue.Name).
			Keyword("exchange", s.exchangeName).
			Keyword("topic", topic).
			Msg("Binding queue")

		if err := s.ch.QueueBind(
			queue.Name,
			topic,
			s.exchangeName,
			false,
			nil,
		); err != nil {
			s.logger.Error(err).
				Keyword("queue_name", queue.Name).
				Keyword("exchange", s.exchangeName).
				Keyword("topic", topic).
				Msg("Failed to binding queue")
		}
	}

	return s.ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}

func (s *Subscriber) RegisterHandler(handler ...MessageHandler) {
	for _, messageHandler := range handler {
		s.handlerList = append(s.handlerList, messageHandler)
	}
}

type SubscriberManagement struct {
	errCh          chan error
	subscriberList []Subscriber
}

func NewSubscriberManagement() (*SubscriberManagement, error) {
	return &SubscriberManagement{
		make(chan error),
		[]Subscriber{},
	}, nil
}

func (s *SubscriberManagement) Register(subscriber Subscriber) {
	s.subscriberList = append(s.subscriberList, subscriber)
}

func (s *SubscriberManagement) Serve() error {
	for _, subscriber := range s.subscriberList {
		go subscriber.Listen()
	}

	return <-s.errCh
}

func (s *SubscriberManagement) GracefulShutdown() error {
	for _, subscriber := range s.subscriberList {
		if err := subscriber.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SubscriberManagement) GetErrCh() chan error {
	return s.errCh
}
