package piazza

import (
	//"bytes"
	//"errors"
	"fmt"
	//"github.com/mpgerlek/piazza-simulator/piazza"
	//"io/ioutil"
	"time"
	"testing"
)


// zookeeper-server-start.sh ~/Downloads/kafka_2.10-0.8.2.0/config/zookeeper.properties
// kafka-server-start.sh ~/Downloads/kafka_2.10-0.8.2.0/config/server.properties
//
// kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
// kafka-topics.sh --list --zookeeper localhost:2181
//
// kafka-console-producer.sh --broker-list localhost:9092 --topic test
// kafka-console-consumer.sh --zookeeper localhost:2181 --topic test --from-beginning

var numRead int

const kafkaHost = "localhost:9092"

func doReads(t *testing.T, numReads *int) {
	r := NewReader(kafkaHost)

	defer func() {
		if err := r.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	partitionConsumer, err := r.ConsumePartition("test3", 0, OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	for  {
		msg := <-partitionConsumer.Messages()
		t.Logf("Consumed: offset:%d  value:%v", msg.Offset, string(msg.Value))
		*numReads++
	}

	t.Log("Reader done: %d", *numReads)
}


func doWrites(t *testing.T, id int, count int) {
	/*topics := kafka.GetTopics()
log.Println(topics)
kafka.AddTopic("foobar")
topics = kafka.GetTopics()
log.Println(topics)*/

	w := NewWriter(kafkaHost)

	defer func() {
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	// TODO: handle "err := <-w.Errors():"

	for n := 0; n<count; n++ {
		w.Input() <- NewKafkaMessage("test3", fmt.Sprintf("mssg %d from %d", n, id))
	}

	t.Logf("Writer done: %d", count)
}


func TestKafka(t *testing.T) {
	var numReads1, numReads2 int

	go doReads(t, &numReads1)
	go doReads(t, &numReads2)

	n := 3
	go doWrites(t, 1, n)
	go doWrites(t, 2, n)

	time.Sleep(1 * time.Second)

	t.Log(numReads1,"---")
	t.Log(numReads2,"---")

	if numReads1 != n*2 {
		t.Fatalf("read1 count was %d, expected %d", numReads1, n*2)
	}
	if numReads2 != n*2 {
		t.Fatalf("read2 count was %d, expected %d", numReads2, n*2)
	}
}
