package cmd

const (
	UserMessage   = iota //玩家发送的消息
	UserBoardCast        //玩家发送的广播
	Kicking              //强制踢人
	Leave                //客户端主动断开连接
)
