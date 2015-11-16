// these are only used for testing purposes

package piazza

import (
	"fmt"
	//"log"
	"strconv"
	"errors"
)

var currPort = 12300

func GetRandomHost() string {
	host := fmt.Sprintf("localhost:%d", currPort)
	currPort++
	return host
}

type PiazzaId string // TODO: this has to be a string, since json.Marshal() only accepts strings as map keys
var BadId PiazzaId = "0"

var nextId int64 = 1 // we skip 0

func NewId() PiazzaId {
	id := nextId
	nextId++
	s := fmt.Sprintf("%d", id)
	pid := PiazzaId(s)
	return pid
}

func NewIdFromString(sid string) (PiazzaId, error) {
	i64, err := strconv.ParseInt(sid, 10, 0)
	if err != nil || i64 < 1 {
		return PiazzaId(0), err
	}
	return NewIdFromInt64(i64)
}

func NewIdFromInt64(i64 int64) (PiazzaId, error) {
	if i64 < 1 {
		return PiazzaId(0), errors.New(fmt.Sprintf("invalid id: %v", i64))
	}
	pid := PiazzaId(strconv.FormatInt(i64, 10))
	return pid, nil
}

func (pid PiazzaId) Int64() (int64, error) {
	i64, err := strconv.ParseInt(pid.String(), 10, 0)
	if err != nil || i64 < 1 {
		return 0, err
	}
	return i64, nil
}


func (id PiazzaId) String() string {
	return string(id)
}
