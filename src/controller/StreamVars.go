package controller

import "github.com/gorilla/websocket"

var (
	upgrader = websocket.Upgrader{}
	commands = make(map[string]chan string)
	outputs  = make(map[string]chan string)
	clients  = make(map[string]map[*websocket.Conn]bool)
)
