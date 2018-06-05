package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gopkg.in/olivere/elastic.v5"

	"github.com/Shopify/sarama"
	"github.com/aarjan/kafka/model"
)

// Service wraps the kafka producer and consumer, and ES client
type Service struct {
	AccessLogProducer sarama.AsyncProducer
	Consumer          sarama.Consumer
	Client            *elastic.Client
}

// CollectData handles the request at endpoint '/'.
// It then produces the events to Kafka topic.
func (s *Service) CollectData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		entry := &model.AccessLogEntry{}

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
func (s *Service) Consume() {
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
				err := s.sendToElasticsearch(msg.Value)
				if err != nil {
					log.Println("error:", err)
				}
				msgCount++

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

// sendToElasticsearch sends event logs to ES
func (s *Service) sendToElasticsearch(body []byte) error {
	ctx := context.Background()

	// Index a accesslog (using JSON serialization)
	put, err := s.Client.Index().
		Index("accesslog").
		Type("log").
		BodyString(string(body)).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed log %s to index %s, type %s\n", put.Id, put.Index, put.Type)

	// Flush to make sure the documents got written.
	_, err = s.Client.Flush().Index("accesslog").Do(ctx)
	if err != nil {
		panic(err)
	}
	return nil
}

// Close closes the connection
func (s *Service) Close() {
	if err := s.AccessLogProducer.Close(); err != nil {
		log.Println("Failed to shutdown log collecter cleanly:", err)
	}
	if err := s.Consumer.Close(); err != nil {
		log.Println("Failed to shutdown consumer cleanly:", err)
	}
}
