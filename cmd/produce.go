package cmd

import (
	"github.com/aarjan/kafka/config"
	"github.com/aarjan/kafka/producer"
	"github.com/aarjan/kafka/server"
	"github.com/aarjan/kafka/service"
)

// var produceCmd = flag.NewFlagSet("produce", flag.ExitOnError)

func produce(c config.AppConfig) {
	// flags
	// br := produceCmd.String("broker", "", "The kafka broker to connect with. Specify a list of brokers separated by ','.")
	// host := produceCmd.String("listen_host", "", "The server host address")
	// port := produceCmd.String("listen_port", "", "The server port address")
	// produceCmd.Parse(os.Args[2:])

	// if *br == "" || *host == "" || *port == "" {
	// 	fmt.Fprintf(os.Stdout, "Please supply the required flags:\n\n")
	// 	produceCmd.Usage()
	// 	os.Exit(1)
	// }
	// brokerList := strings.Split(*br, ",")
	// portAddr, _ := strconv.Atoi(*port)

	app := &service.Service{
		AccessLogProducer: producer.NewProducer(c.Brokers),
	}
	defer app.Close()

	server := server.Server{
		ListenHost: c.ListenHost,
		ListenPort: c.ListenPort,
	}

	// run the producer server
	server.Start(app.CollectData())
}
