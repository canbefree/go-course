package main

import (
	"demo/models/collection"
	"demo/models/user"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

//跨域操作
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	// new Chan ?
	input := make(chan string)
	output := make(chan string)

	coll := collection.NeCollection()
	go coll.Handle(input)

	//用户使用http 方式登陆获取accesstoken
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/views/index.html", 301)
	})

	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("./views/"))))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		uid, err := getUID(r)
		if err != nil {
			w.Write([]byte("sorry,uid must need"))
			return
		}

		var wg sync.WaitGroup

		conn, err := upgrader.Upgrade(w, r, nil)
		defer conn.Close()
		if err != nil {
		}
		log.Printf("uid:%v  connected", uid)
		User := user.NewUser(uid)
		User.Input = input
		User.Output = output

		coll.Join(*User)

		wg.Add(1)

		//处理玩家的 群消息 指定用户私聊消息 玩法和数据
		// go handle(&wg, conn, User)
		go User.Handle(&wg, conn)

		//必须有这个，不然连接立马结束了。
		wg.Wait()
	})

	http.ListenAndServe("localhost:8080", nil)
}

func getUID(r *http.Request) (int, error) {
	uidArr := r.Form["uid"]
	uidStr := uidArr[0]
	return strconv.Atoi(uidStr)
}
