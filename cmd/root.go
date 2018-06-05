package cmd

import (
	"fmt"
	"os"

	"github.com/aarjan/kafka/config"
)

// Execute executes the program
func Execute(c config.AppConfig) {

	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {

	case "consume":
		consume(c)

	case "produce":
		produce(c)

	default:
		fmt.Fprintf(os.Stdout, "%q is not a valid command.\n\n", os.Args[1])
		usage()
	}

}

func usage() {
	fmt.Fprintf(os.Stdout, "Usage:\n")
	fmt.Fprintf(os.Stdout, "  app [command]\n\n")
	fmt.Fprintf(os.Stdout, "Available Commands:\n")
	fmt.Fprintf(os.Stdout, "  produce\tProduce message to Kafka\n")
	fmt.Fprintf(os.Stdout, "  consume\tConsume message from Kafka\n")
}
