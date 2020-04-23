package services

import (
	"xgame/models/cmd"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
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
	cmd_type := gjson.Get(string(msg), "cmd").String()
	switch cmd_type {
	case cmd.CMD_LOGIN:
		logrus.Info("玩家登陆操作:")
		Login.hande
		break
	case cmd.CMD_GUESS:
		break
	default:
		break
	}
}
