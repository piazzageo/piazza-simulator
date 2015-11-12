// Package stringutil contains utility functions for working with strings.
package piazza

import (
	"github.com/Shopify/sarama"
)

const OffsetNewest int64 = sarama.OffsetNewest

type Writer struct {
	producer sarama.AsyncProducer
}

type Reader struct {
	consumer sarama.Consumer
}

func NewKafkaMessage(topic string, data string) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(data)}
}

func NewReader(kafkaHost string) *Reader {
	r := new(Reader)

	consumer, err := sarama.NewConsumer([]string{kafkaHost}, nil)
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

func NewWriter(kafkaHost string) *Writer {

	w := new(Writer)

	config := sarama.NewConfig()
	//config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer([]string{kafkaHost}, config)
	if err != nil {
		panic(err)
	}

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

// this happens aynchronously, so calling GetTopics() immediately afterwards
// will likely not show you your new topic
func AddTopic(kafkaHost string, topic string) {
	broker := sarama.NewBroker(kafkaHost)
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}

	request := sarama.MetadataRequest{Topics: []string{topic}}
	_, err = broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		panic(err)
	}

	if err = broker.Close(); err != nil {
		panic(err)
	}
}

func GetTopics(kafkaHost string) []string {
	broker := sarama.NewBroker(kafkaHost)
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}

	request := sarama.MetadataRequest{ /*Topics: []string{"abba"}*/ }
	response, err := broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		panic(err)
	}

	//topics := make([]string, len(response.Topics))
	topics := []string{}

	for _, v := range response.Topics {
		topics = append(topics, v.Name)
	}

	if err = broker.Close(); err != nil {
		panic(err)
	}

	return topics
}
