package controller

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

// podria enviar simplemente un string con el comando
func GetCommands(c echo.Context) error {

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationMsgpack)
	id := c.Request().Header.Get("id")

	_, exist := commands[id]

	if !exist {
		commands[id] = make(chan string)
	}
	for {
		command := <-commands[id]

		if _, err := fmt.Fprintln(c.Response(), command); err != nil {
			delete(commands, id)
			return err
		}
		time.Sleep(time.Second)
		c.Response().Flush()
	}

}
