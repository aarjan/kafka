package main

import (
	"flag"
	"fmt"
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

	br := consumeCmd.String("broker", "", "The kafka brokers to connect with")
	br2 := produceCmd.String("broker", "", "The kafka brokers to connect with")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "'consume' or 'produce' sub command is required!")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "consume":

		consumeCmd.Parse(os.Args[2:])
		if *br == "" {
			consumeCmd.Usage()
			os.Exit(1)
		}

		brokerList := strings.Split(*br, ",")
		s := &server.Server{
			Consumer: consumer.NewConsumer(brokerList),
			Client:   consumer.NewClient("accesslog"),
		}
		defer s.Close()

		// run the consumer server
		s.Consume()

	case "produce":

		produceCmd.Parse(os.Args[2:])
		if *br2 == "" {
			produceCmd.Usage()
			os.Exit(1)
		}

		brokerList := strings.Split(*br2, ",")
		s := &server.Server{
			AccessLogProducer: producer.NewProducer(brokerList),
		}
		defer s.Close()

		// run the producer server
		panic(http.ListenAndServe(":8080", s.CollectData()))

	default:
		fmt.Fprintln(os.Stderr, "'consume' or 'produce' sub command is required!")
		consumeCmd.Usage()
	}

}
