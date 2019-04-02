package api

import (
	"fmt"
	"net/http"
)

// Init 初始化路由
func init() {

	fmt.Println("route: Init")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	//扩展Writer写法
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
	})

}
