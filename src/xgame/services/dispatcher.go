package services

import (
	"github.com/gorilla/websocket"
)

// type Dispatcher struct {
// 	RoomID int
// 	UserID int
// }

/**
 * 生成全局唯一的房间号
 */
func genarateRoomID() int {
	return 123123
}

// 导出
func HandleRequest(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err == nil {
			switch msgType {
			case websocket.TextMessage:
				DecodeMsg(msg)
			case websocket.CloseMessage:
				break
			default:
			}
		} else {
			break
		}
	}
}

func DecodeMsg(msg []byte) {
	msgType,err := strconv.Atoi(gjson.Get(string(msg),"cmd")	.string
}
