package consumer

import (
	"github.com/Shopify/sarama"
)

// NewConsumer returns a new consumer
func NewConsumer(brokerList []string) sarama.Consumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		panic(err)
	}
	return master
}
