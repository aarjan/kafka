package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aarjan/kafka/consumer"
	"github.com/aarjan/kafka/server"
)

var consumeCmd = flag.NewFlagSet("consume", flag.ExitOnError)

func consume() {
	br := consumeCmd.String("broker", "", "The kafka broker to connect with. Specify a list of brokers separated by ','.")
	consumeCmd.Parse(os.Args[2:])

	if *br == "" {
		fmt.Fprintf(os.Stdout, "Please supply the required flags:\n\n")
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
}
