package test

import (
	"demo/models/protocol"
	"testing"
)

func TestEncode(t *testing.T) {
	client := &protocol.Client{
		CMD:     1,
		Content: "{FID:1,MSG:\"Helo?\"}",
	}
	jsonByte, _ := client.Encode()
	t.Logf("%v", string(jsonByte))
}

func TestDecode(t *testing.T) {
	msg := `
	{"CMD":1,"BODY":"{FID:1,MSG:\"Helo?\"}"}
	`
	client := new(protocol.Client)
	client.Decode([]byte(msg))
	t.Logf("%v", client)
}
