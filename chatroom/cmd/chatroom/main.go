package main

import (
	"chatroom/global"
	"chatroom/server"
	"fmt"
	"log"
	"net/http"
)

const (
	addr = ":2022"
)

func init() {
	global.Init()
}

func main() {
	fmt.Printf("ChatRoom，start on: %s", addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
