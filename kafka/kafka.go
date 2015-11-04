// Package stringutil contains utility functions for working with strings.
package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

type Writer struct {
	producer sarama.AsyncProducer
}

type Reader struct {
	consumer sarama.Consumer
}

func NewReader() *Reader {
	r := new(Reader)

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	r.consumer = consumer

	return r
}

func (r *Reader) Topics() (strs []string, err error) {
	return r.consumer.Topics()
}

func (r *Reader) Partitions(topic string) ([]int32, error) {
	return r.consumer.Partitions(topic)
}

func (r *Reader) ConsumePartition(topic string, partition int32, offset int64) (sarama.PartitionConsumer, error) {
	return r.consumer.ConsumePartition(topic, partition, offset)
}

func (r *Reader) Close() error {
	return r.consumer.Close()
}

func NewWriter() *Writer {
	w := new(Writer)

	config := sarama.NewConfig()
	//config.Producer.Return.Successes = true

	log.Println("PRE")
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	log.Println("POST")

	w.producer = producer
	return w
}

func (w *Writer) Close() error {
	return w.producer.Close()
}

func (w *Writer) Input() chan<- *sarama.ProducerMessage {
	return w.producer.Input()
}

func (w *Writer) Successes() <-chan *sarama.ProducerMessage {
	return w.producer.Successes()
}

func (w *Writer) Errors() <-chan *sarama.ProducerError {
	return w.producer.Errors()
}
