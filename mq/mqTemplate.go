package mq

import "github.com/streadway/amqp"

// Producer rabbitMq生产者
type Producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 交换机
	exchange string
	// key
	key string
	// mq的链接
	MqUrl string
	// 交换机类型
	exchangeType string
	// 队列名称
	queueName string
	// 是否持久化
	durable bool
	// 是否自动删除
	autoDelete bool
}

// NewProducer 创建一个生产者
func NewProducer(mqUrl, exchange, exchangeType, queueName, key string, durable, autoDelete bool) (*Producer, error) {
	producer := &Producer{
		MqUrl:        mqUrl,
		exchange:     exchange,
		key:          key,
		exchangeType: exchangeType,
		queueName:    queueName,
		durable:      durable,
		autoDelete:   autoDelete,
	}
	var err error
	producer.conn, err = amqp.Dial(producer.MqUrl)
	if err != nil {
		return nil, err
	}
	producer.channel, err = producer.conn.Channel()
	if err != nil {
		return nil, err
	}
	err = producer.channel.ExchangeDeclare(
		producer.exchange,
		producer.exchangeType,
		producer.durable,
		producer.autoDelete,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	_, err = producer.channel.QueueDeclare(
		producer.queueName,
		producer.durable,
		producer.autoDelete,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = producer.channel.QueueBind(
		producer.queueName,
		producer.key,
		producer.exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// PublishSimple 发送消息
func (p *Producer) PublishSimple(message string) error {
	return p.channel.Publish(
		p.exchange,
		p.key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

// Close 关闭链接
func (p *Producer) Close() error {
	err := p.channel.Close()
	if err != nil {
		return err
	}
	err = p.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

// Consumer 消费者
type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 交换机
	exchange string
	// key
	key string
	// mq的链接
	MqUrl string
	// 交换机类型
	exchangeType string
	// 队列名称
	queueName string
}

// NewConsumer 创建一个消费者
func NewConsumer(mqUrl, exchange, exchangeType, queueName, key string) (*Consumer, error) {
	consumer := &Consumer{
		MqUrl:        mqUrl,
		exchange:     exchange,
		key:          key,
		exchangeType: exchangeType,
		queueName:    queueName,
	}
	var err error
	consumer.conn, err = amqp.Dial(consumer.MqUrl)
	if err != nil {
		return nil, err
	}
	consumer.channel, err = consumer.conn.Channel()
	if err != nil {
		return nil, err
	}
	err = consumer.channel.ExchangeDeclare(
		consumer.exchange,
		consumer.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	_, err = consumer.channel.QueueDeclare(
		consumer.queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = consumer.channel.QueueBind(
		consumer.queueName,
		consumer.key,
		consumer.exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// ConsumeSimple 消费消息
func (c *Consumer) ConsumeSimple() (<-chan amqp.Delivery, error) {
	msg, err := c.channel.Consume(
		c.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
