package piazza

import (
	"strings"
	"testing"
	"bytes"
)

func TestMessage(t *testing.T) {
	// Object to Object's JSON
	m := NewMessage()
	m.Id = 10
	m.Type = CreateJobMessage
	m.User = User{Id: 123, UserType: NormalUser, Name: "Bob"}
	m.Status = SuccessStatus

	params := map[string]Any{"A": "alpha", "B": 222}
	m.CreateJobPayload = CreateJobPayload{ServiceName: "echo", Parameters: params, Comment: "Hi."}

	jsonActual, err := m.ToJson()
	if err != nil {
		t.Error(err)
	}

	buf, err := m.ToBytes()
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(buf, []byte(jsonActual)) != 0 {
		t.Error("bytes: actual != expected")
	}

	// JSON to Object's JSON
	jsonExpected :=
		`{
	"id":10,
	"type":1,
	"timestamp":"0001-01-01T00:00:00Z",
	"user":{"Id":123,"UserType":1,"Name":"Bob"},
	"status":1,
	"error":null,
	"percentComplete":0,
	"timeRemaining":0,
	"service": {"id":"", "name":"", "Type":0, "description":""}
}`
	jsonExpected = strings.Replace(jsonExpected, " ", "", -1)
	jsonExpected = strings.Replace(jsonExpected, "\t", "", -1)
	jsonExpected = strings.Replace(jsonExpected, "\n", "", -1)

	if strings.Compare(jsonActual[:3], jsonExpected[:3]) != 0 {
		t.Logf("actual: (%d) %s", len(jsonActual), jsonActual)
		t.Logf("expected: (%d) %s", len(jsonExpected), jsonExpected)
		t.Errorf("json: actual != expected")
	}

	// JSON to Object
	m2, err := NewMessageFromBytes([]byte(jsonActual))
	if err != nil {
		t.Error(err)
	}

	if m2.Id != m.Id {
		t.Error("actual field != expected field")
	}
	if 	m2.User.Name != m.User.Name  {
		t.Error("actual field != expected field")
	}
	sA := m2.CreateJobPayload.Parameters["A"].(string)
	sE := m.CreateJobPayload.Parameters["A"].(string)
	if strings.Compare(sA, sE) != 0 {
		t.Error("actual field != expected field")
	}
	var iA float64 = m2.CreateJobPayload.Parameters["B"].(float64) // json makes everything a float64
	var iE int = m.CreateJobPayload.Parameters["B"].(int) // but we manually put an int into the "any"
	if iA != float64(iE) {
		t.Error("actual field != expected field")
	}
	t.Log(iA, iE)
}
