# Go Kafka ES client

A go CLI that,

- listens to an event through a HTTP server

- produces messages to a Kafka topic

- consumes messages from a Kafka topic

- sends the messages to a Elasticsearch server

## Requirements

- Running instance of Kafka on usual ports

        - For simplicity, you can run the docker container
        ```shell
            sudo docker run --rm --net=host landoop/fast-data-dev
        ```

- Running instance of Elasticsearch on usual ports

## Run a HTTP server and produce event to Kafka

```shell
    ./kafka produce --broker=localhost:9092 --listen_host=localhost --listen_port=8080
```

## Consume message from Kafka server and send it to Elasticsearch

```shell
    ./kafka consume --broker=localhost:9092 --es_host=localhost -es_port=9200 --es_index=access_log
```