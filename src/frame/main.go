package main

import "flag"

// go -N 12 -C 23 url

var (
	count       string
	concurrence string
	src 		string
)

func main() {
	flag.StringVar(&count, "N", "1", "")
	flag.StringVar(&concurrence, "C", "123", "")
	flag.StringVar(&src, "S", "123", "")
	flag.Parse()

}
