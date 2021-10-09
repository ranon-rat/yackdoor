package controller

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func UploadCommand(c echo.Context) error {

	id := c.QueryParam("id")

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	if _, exist := clients[id]; !exist {

		clients[id] = make(map[*websocket.Conn]bool)
	}

	clients[id][ws] = true

	if _, exist := commands[id]; !exist {
		fmt.Println("Command not found")
		ws.WriteMessage(websocket.TextMessage, []byte(`{"output":"sorry not avaible"}`))

		return nil
	}
	defer ws.Close()
	go func() {
		for {

			output := <-outputs[id]

			fmt.Println(output)
			for c := range clients[id] {
				fmt.Println(c.RemoteAddr().String())
				fmt.Println(len(clients[id]))
				if err := c.WriteMessage(websocket.TextMessage, []byte(output)); err != nil { // if the socket is closed

					delete(clients[id], c)

					if len(clients[id]) == 0 {
						delete(clients, id)

					}
					return
				}

			}

		}

	}()
	for {

		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			delete(clients[id], ws)
			if len(clients[id]) == 0 {
				delete(clients, id)

			}
			return err
		}

		commands[id] <- string(msg)

	}

}
