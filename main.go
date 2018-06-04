package main

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/aarjan/kafka/producer"

	"github.com/aarjan/kafka/consumer"
	"github.com/aarjan/kafka/server"
)

func main() {

	consumeCmd := flag.NewFlagSet("consume", flag.ExitOnError)
	produceCmd := flag.NewFlagSet("produce", flag.ExitOnError)

	br := flag.String("broker", "localhost:9092", "The kafka brokers to connect with")
	// flag.Parse()
	brokerList := strings.Split(*br, ",")

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "consume":
		consumeCmd.Parse(os.Args[2:])
		s := &server.Server{
			Consumer: consumer.NewConsumer(brokerList),
		}
		defer s.Close()
		s.Consume()
	case "produce":
		produceCmd.Parse(os.Args[2:])
		s := &server.Server{
			AccessLogProducer: producer.NewProducer(brokerList),
		}
		defer s.Close()
		panic(http.ListenAndServe(":8080", s.CollectData()))
	}

}
