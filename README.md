# Go Kafka ES client

A go CLI that,

- listens to an event through a HTTP server

- produces messages to a Kafka topic

- consumes messages from a Kafka topic

- sends the messages to a Elasticsearch server

## Run a HTTP server 

```shell
    ./kafka produce --broker localhost:9200
```

## Consume message from Kafka server and send it to Elasticsearch

```shell
    ./kafka consume --broker localhost:9200
```