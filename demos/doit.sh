#!/bin/sh

./start-zookeeper.sh &
sleep 5

./start-kafka.sh &
sleep 5

./start-elasticsearch.sh &
sleep 5

./start-discover.sh &
sleep 5

./do-register.sh

./start-logger.sh &
sleep 1

./start-uuidgen.sh &
sleep 1
