package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	"github.com/aarjan/kafka/producer"
)

// Server wraps the kafka producer
type Server struct {
	AccessLogProducer sarama.AsyncProducer
	Consumer          sarama.Consumer
}

func (s *Server) CollectData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		entry := &producer.AccessLogEntry{}

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
		log.Printf("sending data to kafka: %v", entry)
		w.WriteHeader(200)
	})
}

// Consume consumes the message and sents it to elasticsearch
func (s *Server) Consume() {
	consumer, err := s.Consumer.ConsumePartition("access_log", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// count the number of msg consumed
	msgCount := 0

	// signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Recieved messages", string(msg.Value))
			case <-signals:
				fmt.Println("\tInterrupt detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	fmt.Println("Processed:", msgCount, "messages")
	os.Exit(1)
}

// Close closes the connection
func (s *Server) Close() {
	if err := s.AccessLogProducer.Close(); err != nil {
		log.Println("Failed to shutdown log collecter cleanly:", err)
	}
	if err := s.Consumer.Close(); err != nil {
		log.Println("Failed to shutdown consumer cleanly:", err)
	}
}
