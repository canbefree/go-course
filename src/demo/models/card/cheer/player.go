package cheer

import "demo/models/user"

//玩家
type Player struct {
	User user.User //属于哪个用户 跟服务器交互
}

type IPlayer interface {
	Join() bool  //玩家加入游戏
	Ready() bool //玩家准备
	Start() bool //房主才有的功能 如果是房主开始游戏 --开始一个线程
}
