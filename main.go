package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	br := flag.String("broker", "localhost:9092", "The kafka brokers to connect with")
	brokerList := strings.Split(*br, ",")
	flag.Parse()

	s := &Server{newProducer(brokerList)}
	panic(http.ListenAndServe(":8080", s.collectData()))
}

// Server wraps the kafka producer
type Server struct {
	AccessLogProducer sarama.AsyncProducer
}

func (s *Server) collectData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		entry := &accessLogEntry{}

		err := json.NewDecoder(r.Body).Decode(entry)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		entry.TimeStamp = time.Now()

		s.AccessLogProducer.Input() <- &sarama.ProducerMessage{
			Topic: "access_log",
			Key:   nil,
			Value: entry,
		}
	})
}

type accessLogEntry struct {
	Msg       string    `json:"msg"`
	Code      int       `json:"code"`
	TimeStamp time.Time `json:"time"`
	encoded   []byte
	err       error
}

func (ale *accessLogEntry) ensureEncoded() {
	if ale.err == nil && ale.encoded == nil {
		ale.encoded, ale.err = json.Marshal(ale)
	}
}

func (ale *accessLogEntry) Length() int {
	ale.ensureEncoded()
	return len(ale.encoded)
}
func (ale *accessLogEntry) Encode() ([]byte, error) {
	ale.ensureEncoded()
	return ale.encoded, ale.err
}

func newProducer(brokerList []string) sarama.AsyncProducer {

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
