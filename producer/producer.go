package producer

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

// AccessLogEntry is a schema for individual log
type AccessLogEntry struct {
	Msg       string    `json:"msg"`
	Code      int       `json:"code"`
	TimeStamp time.Time `json:"time"`
	encoded   []byte
	err       error
}

func (ale *AccessLogEntry) ensureEncoded() {
	if ale.err == nil && ale.encoded == nil {
		ale.encoded, ale.err = json.Marshal(ale)
	}
}

func (ale *AccessLogEntry) Length() int {
	ale.ensureEncoded()
	return len(ale.encoded)
}
func (ale *AccessLogEntry) Encode() ([]byte, error) {
	ale.ensureEncoded()
	return ale.encoded, ale.err
}

// NewProducer returns a new sarama.AsyncProducer
func NewProducer(brokerList []string) sarama.AsyncProducer {

	// For the access log, we are looking for AP semantics, with high throughput.
	// By creating batches of compressed messages, we reduce network I/O at a cost of more latency.
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()

	return producer
}
