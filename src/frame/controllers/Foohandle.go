package controllers

import "net/http"

type FooHandle struct {
}

func (fh FooHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,world"))
}

