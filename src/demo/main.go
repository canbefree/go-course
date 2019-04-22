package main

import (
	"fmt"
	"net/http"
	"sync"

	"demo/models/user"

	"github.com/gorilla/websocket"
)

//跨域操作
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	//用户使用http 方式登陆获取accesstoken
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/views/index.html", 301)
	})

	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("./views/"))))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		accessToken := r.Form["accessToken"]
		if accessToken == nil {
			return
		}

		User := user.NewUser(accessToken[0])

		conn, err := upgrader.Upgrade(w, r, nil)
		defer conn.Close()
		if err != nil {

		}
		var wg sync.WaitGroup
		wg.Add(1)

		//处理玩家的 群消息 指定用户私聊消息 玩法和数据
		go handle(&wg, conn, User)

		//必须有这个，不然连接立马结束了。
		wg.Wait()
	})

	http.ListenAndServe("localhost:8080", nil)
}

func handle(wg *sync.WaitGroup, conn *websocket.Conn, user *user.User) {

	defer wg.Done()
	for {
		msgType, msg, err := conn.ReadMessage()
		if err == nil {
			fmt.Printf("%v:%v:%v", msgType, string(msg), user)
		} else {
			break
		}
	}

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
