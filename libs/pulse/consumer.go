package pulse

import (
	"ms.api/config"
	"ms.api/utils"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pkg/errors"
	"log"
	"time"
)

type pulseConsumer struct {
	binding pulsar.Consumer
}

type Event struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	Data      []utils.JSON `json:"data"`
	CreatedOn time.Time    `json:"created_on"`
}

type SubscriptionHandler func(event Event) (status bool, err error)

type Consumer interface {
	Subscribe(handler SubscriptionHandler)
	GetBinding() interface{}
}

func NewConsumer(client pulsar.Client, topic string) (Consumer, error) {
	if client == nil {
		return nil, errors.New("broker client is dead")
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:                       topic,
		AutoDiscoveryPeriod:         0,
		SubscriptionName:            config.ServiceName,
		Type:                        pulsar.Shared,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionLatest,
		Name:                        config.ServiceName,
	})

	if err != nil {
		log.Println(err, " Failed to create consumer")
		return nil, err
	}

	return &pulseConsumer{binding: consumer}, nil

}

func (c *pulseConsumer) Subscribe(handler SubscriptionHandler) {
	for {
		if val, ok := <-c.binding.Chan(); ok {
			go func(cm pulsar.ConsumerMessage) {
				val.Message.Topic()
				val.Message.EventTime()
				id := utils.ByteToHex(val.Message.ID().Serialize())
				// TODO: Ensure event struct is according to the Roava Ecosystem.
				event := Event{
					Id:        id,
					Name:      val.Message.Topic(),
					CreatedOn: val.Message.EventTime(),
				}
				var data []utils.JSON
				_ = json.Unmarshal(val.Message.Payload(), &data)
				event.Data = data

				// TODO: Agree on what we want to do to panicked handler.
				ok, err := handler(event)
				if err != nil {
					log.Printf("Subscription handler failed to process %s failed with error: %s", cm.Message.ID(), err)
				}
				if ok {
					cm.AckID(cm.Message.ID())
				}
			}(val)
		}
	}
}

func (c *pulseConsumer) GetBinding() interface{} {
	return c.binding
}
