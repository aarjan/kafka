package cmd

import (
	"fmt"

	"github.com/aarjan/kafka/config"
	"github.com/aarjan/kafka/consumer"
	"github.com/aarjan/kafka/service"
)

// var consumeCmd = flag.NewFlagSet("consume", flag.ExitOnError)

func consume(c config.AppConfig) {
	// flags
	// br := consumeCmd.String("broker", "", "The kafka broker to connect with. Specify a list of brokers separated by ','.")
	// host := consumeCmd.String("es_host", "localhost", "The elasticsearch host address")
	// port := consumeCmd.String("es_port", "9200", "The elasticsearch port address")
	// index := consumeCmd.String("es_index", "", "The elasticsearch index to post the event logs")
	// consumeCmd.Parse(os.Args[2:])

	// if *br == "" || *index == "" {
	// 	fmt.Fprintf(os.Stdout, "Please supply the required flags:\n\n")
	// 	consumeCmd.Usage()
	// 	os.Exit(1)
	// }
	// brokerList := strings.Split(*br, ",")
	address := fmt.Sprintf("%s:%s", c.ESHost, c.ESPort)

	s := &service.Service{
		Consumer: consumer.NewConsumer(c.Brokers),
		Client:   consumer.NewClient(address, c.Index),
	}
	defer s.Close()

	// run the consumer service
	s.Consume()
}
