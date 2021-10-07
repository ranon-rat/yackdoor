package controller

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadCommand(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	id := c.QueryParam("id")
	buf := new(bytes.Buffer)

	buf.ReadFrom(c.Request().Body)
	asd := buf.String()

	if _, exist := commands[id]; !exist {
		c.String(http.StatusNotFound, "sorry not avaible")
	}
	commands[id] <- asd

	for {
		output := <-outputs[id]
		if _, err := io.Copy(c.Response(), bytes.NewBufferString(output)); err != nil {
			commands[id] <- "break"
			delete(outputs, id)
			return err
		}

		c.Response().Flush()
	}

}
