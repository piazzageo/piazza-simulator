package jobs

import (
	"fmt"
	"log"
	"time"
)

type MessageId int64

var currentId MessageId = 1

type MessageType int

const (
	JobRequest    MessageType = 1
	StatusRequest MessageType = 2
)

type Message struct {
	id        MessageId
	mtype     MessageType
	timestamp time.Time
}

func NewMessage(mtype MessageType) *Message {
	var m = Message{id: currentId, timestamp: time.Now(), mtype: mtype}
	currentId++

	return &m
}

func (m Message) String() string {
	return fmt.Sprintf("{id:%v mtype:%v}", m.id, m.mtype)
}
