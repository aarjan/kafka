# Go Kafka ES client

A go CLI that,

- listens to an event through a HTTP server at endpoint '/event'

- produces messages to a Kafka topic

- consumes messages from a Kafka topic

- sends the messages to a Elasticsearch server

## Requirements

- Running instance of Kafka on usual ports
  (for simplicity, you can run the docker container)

``` shell
    sudo docker run --rm --net=host landoop/fast-data-dev
```

- Running instance of Elasticsearch on usual ports

## To Run the application

- Make a new file __run.env__ & copy the contents of [run.env.sample](run.env.sample) to run.env.
  Make sure, you supply the required environment varibles in run.env.

  You need to change the run.sh to run two separate servers for Kafka producers and Kafka consumers

- __Run a HTTP server and produce event to Kafka__

```shell
    ./kafka produce
```

- __Consume message from Kafka server and send it to Elasticsearch__

```shell
    ./kafka consume
```

### Build & run the program

``` shell
    ./dev.sh
```

You can run [test_messages](test-msg-producer/test_messages.go) to send dummy message in running Kafka producer server.