package main

import (
	"github.com/mpgerlek/piazza-simulator/kafka"
	"log"
	"os"
	"os/signal"
)

func main() {

	r := kafka.NewReader()

	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := r.ConsumePartition("test3", 0, kafka.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			log.Printf("   Value %v\n", string(msg.Value))
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
