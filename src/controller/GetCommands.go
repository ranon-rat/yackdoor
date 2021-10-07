package controller

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// YOU ONLY CAN USE THIS WITH ONLY 1 CLIENT

func GetCommands(c echo.Context) error {

	id := c.Request().Header.Get("id") //later i will change this to the headers but for now it works
	// its just for see the id of the client

	if _, exist := commands[id]; !exist {
		commands[id] = make(chan string)
		outputs[id] = make(chan string)
	}
	websocket.Handler(func(conn *websocket.Conn) {

		go func() {
			// this is for check if it still connected
			for {
				if err := websocket.Message.Send(conn, "you still there?"); err != nil {
					deleteThisPlease()
					return
				}
				time.Sleep(time.Minute * 10)
			}
		}()
		go func() {
			for {

				msg := ""
				if err := websocket.Message.Receive(conn, &msg); err != nil {
					log.Println(msg)

					deleteThisPlease()
					return
				}

				outputs[id] <- msg
			}
		}()
		for {
			command := <-commands[id]
			log.Println(command)
			if err := websocket.Message.Send(conn, command); err != nil {
				deleteThisPlease()
				return
			}

		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func deleteThisPlease() {
	delete(commands, "id")
	delete(outputs, "id")
}
