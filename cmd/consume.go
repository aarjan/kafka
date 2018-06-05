package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aarjan/kafka/consumer"
	"github.com/aarjan/kafka/service"
)

var consumeCmd = flag.NewFlagSet("consume", flag.ExitOnError)

func consume() {
	// flags
	br := consumeCmd.String("broker", "", "The kafka broker to connect with. Specify a list of brokers separated by ','.")
	index := consumeCmd.String("index", "", "The elasticsearch index to post the event logs")
	consumeCmd.Parse(os.Args[2:])

	if *br == "" || *index == "" {
		fmt.Fprintf(os.Stdout, "Please supply the required flags:\n\n")
		consumeCmd.Usage()
		os.Exit(1)
	}
	brokerList := strings.Split(*br, ",")

	s := &service.Service{
		Consumer: consumer.NewConsumer(brokerList),
		Client:   consumer.NewClient(*index),
	}
	defer s.Close()

	// run the consumer service
	s.Consume()
}
