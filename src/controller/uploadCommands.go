package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func UploadCommand(c echo.Context) error {

	id := c.QueryParam("id")
	breakTHIS := make(chan bool)
	websocket.Handler(func(conn *websocket.Conn) {
		clients[id] = make(map[*websocket.Conn]bool)
		clients[id][conn] = true

		if _, exist := commands[id]; !exist {
			fmt.Println("Command not found")
			websocket.Message.Send(conn, `{"output":"sorry not avaible"}`)
			conn.Close()
			return
		}
		defer conn.Close()
		go func() {
			for {

				output := <-outputs[id]

				fmt.Println(output)
				for c := range clients[id] {
					if err := websocket.Message.Send(c, output); err != nil {

						commands[id] <- "break"
						delete(outputs, id)
						conn.Close()

						breakTHIS <- true
						return
					}
				}

			}

		}()
		for {

			select {
			case <-breakTHIS:
				conn.Close()
				return

			default:
				msg := ""

				if err := websocket.Message.Receive(conn, &msg); err != nil {
					fmt.Println(err)
					commands[id] <- "break"
					delete(outputs, id)
					conn.Close()
					return

				}

				commands[id] <- msg

			}

		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil

}
