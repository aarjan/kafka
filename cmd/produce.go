package cmd

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aarjan/kafka/producer"
	"github.com/aarjan/kafka/server"
)

var produceCmd = flag.NewFlagSet("produce", flag.ExitOnError)

func produce() {
	br := consumeCmd.String("broker", "", "The kafka broker to connect with. Specify a list of brokers separated by ','.")
	produceCmd.Parse(os.Args[2:])

	if *br == "" {
		fmt.Fprintf(os.Stdout, "Please supply the required flags:\n\n")
		produceCmd.Usage()
		os.Exit(1)
	}
	brokerList := strings.Split(*br, ",")

	s := &server.Server{
		AccessLogProducer: producer.NewProducer(brokerList),
	}
	defer s.Close()

	// run the producer server
	panic(http.ListenAndServe(":8080", s.CollectData()))
}
