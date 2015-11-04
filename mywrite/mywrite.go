// zookeeper-server-start.sh ~/Downloads/kafka_2.10-0.8.2.0/config/zookeeper.properties 
// kafka-server-start.sh ~/Downloads/kafka_2.10-0.8.2.0/config/server.properties
//
// kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
// kafka-topics.sh --list --zookeeper localhost:2181
//
// kafka-console-producer.sh --broker-list localhost:9092 --topic test 
// kafka-console-consumer.sh --zookeeper localhost:2181 --topic test --from-beginning

package main

import (
	"github.com/Shopify/sarama"

	//"crypto/tls"
	//"crypto/x509"
	//"encoding/json"
	//"flag"
	//"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	"os"
	
	"os/signal"
)

func main() {
	config := sarama.NewConfig()
	//config.Producer.Return.Successes = true
	
	log.Println("PRE")
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
	    panic(err)
	}
	log.Println("POST")

	defer func() {
	    if err := producer.Close(); err != nil {
	        log.Fatalln(err)
	    }
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	log.Println("ProducerMessage")

	var enqueued, errors int
	ProducerLoop:
	for {
	    select {
	    case producer.Input() <- &sarama.ProducerMessage{Topic: "test3", Key: nil, Value: sarama.StringEncoder("testing 123")}:
	        enqueued++
	    case err := <-producer.Errors():
	        log.Println("Failed to produce message", err)
	        errors++
	    case <-signals:
	        break ProducerLoop
	    }
		
		if enqueued==10 {
			break
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
