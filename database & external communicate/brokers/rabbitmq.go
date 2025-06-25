package brokers

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	_ "github.com/streadway/amqp"
	"log"
	"time"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchange     = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "test-key", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
	lifetime     = flag.Duration("lifetime", 5*time.Second, "lifetime of process before shutdown (0s=infinite)")
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ() *RabbitMQ {
	log.Printf("dialing %q", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &RabbitMQ{
		conn: conn,
	}
}

func (rabbit *RabbitMQ) Publish(body []byte) {
	log.Printf("got connection, getting channel")
	channel, err := rabbit.conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = channel.ExchangeDeclare(
		"gotr-city-exchange",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		fmt.Println(err)
	}

	if err = channel.Publish(
		"gotr-city-exchange",
		"",
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient,
			Priority:        0,
		},
	); err != nil {
		fmt.Println(err)
	}
}
