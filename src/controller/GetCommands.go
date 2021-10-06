package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// YOU ONLY CAN USE THIS WITH ONLY 1 CLIENT

func GetCommands(c echo.Context) error {
	fmt.Println(c.Request().URL)
	id := c.QueryParam("id")         //later i will change this to the headers but for now it works
	fmt.Println(c.Request().URL, id) // its just for see the id of the client
	_, exist := commands[id]

	if !exist {
		commands[id] = make(chan string)
	}
	websocket.Handler(func(conn *websocket.Conn) {

		go func() {
			// this is for check if it still connected
			for {
				if err := websocket.Message.Send(conn, "you still there?"); err != nil {
					delete(commands, id)
					return
				}
				time.Sleep(time.Minute * 10)
			}
		}()
		go func() {
			for {

				msg := ""

				err := websocket.Message.Receive(conn, &msg)
				if err != nil {
					log.Println(msg)
					delete(commands, id)
					return
				}

				fmt.Println(msg)
			}
		}()
		for {
			command := <-commands[id]
			log.Println(command)
			if err := websocket.Message.Send(conn, command); err != nil {
				delete(commands, id)
				return
			}

		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
