#!/bin/sh

kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test

kafka-topics.sh --list --zookeeper localhost:2181

kafka-console-producer.sh --broker-list localhost:9092 --topic test

kafka-console-consumer.sh --zookeeper localhost:2181 --topic test --from-beginning

