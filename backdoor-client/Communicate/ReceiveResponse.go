package Communicate

import "golang.org/x/net/websocket"

func StartEverything(conn *websocket.Conn) {

	go ExecuteCommand()
	go KillProcess()
	go GetOutAndErr(conn)
	GetCommand(conn)
}
