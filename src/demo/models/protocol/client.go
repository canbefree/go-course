package protocol

import "encoding/json"

type Client struct {
	Base
	Func string //回调客户端
}

func (p *Client) Decode(msg []byte) {
	json.Unmarshal(msg, p)
}

func (p *Client) Encode() ([]byte, error) {
	return json.Marshal(p)
}
