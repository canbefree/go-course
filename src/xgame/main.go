package main

import (
	"net/http"
	"xgame/services"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

//跨域操作
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", ws)
	http.HandleFunc("/boardcast",boardCast)
	http.HandleFunc("/sendMsg",sendMsg)
	http.ListenAndServe("localhost:8888", nil)
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err == nil {
		logrus.Info("协议升级：")
		go services.HandleRequest(conn)
	}
}




//全员推送广播
func boardCast(){
	
}

//单独给某人推送消息
func sendMsg(){

}