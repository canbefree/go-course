package main

import (
	"fmt"
	"frame/models"
)

func main() {
	u := models.User{"12", "小明", "我的"}
	fmt.Printf(u.Name())
	fmt.Printf("end")

	ulist := make(map[string][]models.User)

	var ms = []models.User{u, u, u}

	ulist["hah"] = ms

	fmt.Println(ulist)
}

