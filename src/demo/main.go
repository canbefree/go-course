package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

//跨域操作
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,world"))
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {

		}
		//新建一个线程 这个线程要处理玩家的 群消息 指定用户私聊消息 玩法和数据
		go handle(conn)
	})

	http.ListenAndServe("localhost:8080", nil)
}

func handle(conn) {
	// 新客户端连接

	//订阅room

	//接受历史消息？

	/**
	 *  循环监听客户端的消息
	 * 		1：群消息，广播到群内
	 *		2: 私聊
	 * 		3: Player消息
	 */
	println(conn)
}
