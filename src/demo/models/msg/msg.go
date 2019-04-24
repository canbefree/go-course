package msg

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

const (
	//CMDBoardCast 广播消息
	CMDBoardCast = iota
	//CMDNormal 普通消息
	CMDNormal
)

//Msg 消息
type Msg struct {
	//命令行
	CMD int
	//消息内容
	Body string
	//回应
	Response string
}

func (msg Msg) GetCMD() int {
	return msg.CMD
}

func (msg Msg) GetBody() string {
	return msg.Body
}

func (msg Msg) Handle(conn *websocket.Conn) error {
	conn.WriteMessage(websocket.TextMessage, []byte(msg.Response))
	return errors.New("msg")
}

//JSONDecode 解析客户端发过来的消息
func JSONDecode(orginMsg string) (iMsg, error) {
	cmd := gjson.Get(orginMsg, "CMD").Int()
	log.Println(cmd)
	log.Println(orginMsg)
	err := errors.New("error")
	switch cmd {
	case CMDBoardCast:
		var msgBoardCast BoardCast
		err = json.Unmarshal([]byte(orginMsg), &msgBoardCast)
		return msgBoardCast, err
	case CMDNormal:
		var msgNormal Normal
		err = json.Unmarshal([]byte(orginMsg), &msgNormal)
		return msgNormal, err
	default:
		return nil, err
	}
}
