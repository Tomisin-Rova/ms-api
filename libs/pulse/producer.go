package pulse

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

type pulseProducer struct {
	binding pulsar.Producer
}

type Producer interface {
	Publish(message []byte) error
	GetBinding() interface{}
}

func NewProducer(client pulsar.Client, topic string) (Producer, error) {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		log.Println("Failed to create new producer with the following error", err)
		return nil, err
	}

	return &pulseProducer{binding: producer}, nil

}

// TODO: Tidy our publishing according to the Roava Event structure.
func (p *pulseProducer) Publish(message []byte) error {
	id, e := p.binding.Send(context.Background(), &pulsar.ProducerMessage{
		Payload:   message,
		EventTime: time.Now(),
	})

	if e != nil {
		log.Println(e, " Failed to send message.")
		return e
	}

	log.Println("Pulsar Message Sent: ", id)
	return nil
}

func (p *pulseProducer) GetBinding() interface{} {
	return p.binding
}
