package controller

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadCommand(c echo.Context) error {
	id := c.QueryParam("id")
	buf := new(bytes.Buffer)

	buf.ReadFrom(c.Request().Body)
	asd := buf.String()
	_, exist := commands[id]
	if !exist {
		c.String(http.StatusNotFound, "sorry not avaible")
	}
	commands[id] <- asd
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationMsgpack)
	for {
		output := <-outputs[id]

		if _, err := io.Copy(c.Response(), bytes.NewBufferString(output)); err != nil {
			commands[id] <- "break"
			return err
		}

		c.Response().Flush()
	}

}
