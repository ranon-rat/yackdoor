package controller

import "golang.org/x/net/websocket"

var (
	commands = make(map[string]chan string)
	outputs  = make(map[string]chan string)
	clients  = make(map[string]map[*websocket.Conn]bool)
)
