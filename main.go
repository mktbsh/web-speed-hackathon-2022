package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/mktbsh/web-speed-hackathon-2022/docs"
)

// @title CyberTicket API
// @version 1.0
// @description CyberTicketサーバのドキュメント
// @host localhost:1323
// @BasePath /
func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
