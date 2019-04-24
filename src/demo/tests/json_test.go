package test

import (
	"demo/models/msg"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/tidwall/gjson"
)

func TestGjsonDecode(t *testing.T) {
	subject := "{\"CMD\":12,\"Msg\":\"Red\"}"
	except := "Red"
	result := gjson.Get(subject, "Msg").String()
	if except != result {
		t.Error("error")
	}

}

func TestJsonEecode(t *testing.T) {
	message := msg.Msg{CMD: 12, Body: "Red"}
	strByte, err := json.Marshal(message)
	if err != nil {
		t.Error(err)
	}

	strString := string(strByte)

	except := "{\"CMD\":12,\"Msg\":\"Red\"}"

	if strString != except {
		t.Errorf("result:%v  jt:%v", strString, except)
	}

	// os.Stdout.Write(jt)
	t.Logf("byte:%v", strByte)
	t.Logf("byte:%v", []byte(except))
}

func TestJsonDecode(t *testing.T) {
	subject := "{\"CMD\":12,\"Msg\":\"Red\"}"
	var message msg.Msg
	//注意这里必须传 指针
	err := json.Unmarshal([]byte(subject), &message)
	if err != nil {
		t.Error("error")
	}
	except := msg.Msg{CMD: 12, Body: "Red"}
	if except != message {
		t.Errorf("not equal %v != %v", except, message)
	}
	t.Log(message)
}

func TestJsonDecoder(t *testing.T) {
	//读取文件流
	const jsonStream = `
        {"Name": "Ed", "Text": "Knock knock."}
        {"Name": "Sam", "Text": "Who's there?"}
        {"Name": "Ed", "Text": "Go fmt."}
        {"Name": "Sam", "Text": "Go fmt who?"}
        {"Name": "Ed", "Text": "Go fmt yourself!"}
    `
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

func TestJsonEncoder(t *testing.T) {
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func BenchmarkGjsonPerformance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subject := "{\"CMD\":12,\"Msg\":\"Red\"}"
		_ = gjson.Get(subject, "Msg").String()
	}
}

func BenchmarkJsonPerformance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subject := "{\"CMD\":12,\"Msg\":\"Red\"}"
		var message msg.Msg
		_ = json.Unmarshal([]byte(subject), &message)
	}
}
