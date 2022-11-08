package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/mktbsh/web-speed-hackathon-2022/docs"
)

//go:embed all:dist
var frontendFS embed.FS

func buildFrontendFS() http.FileSystem {
	build, err := fs.Sub(frontendFS, "dist")
	if err != nil {
		log.Fatal(err)
	}

	return http.FS(build)
}

// @title CyberTicket API
// @version 1.0
// @description CyberTicketサーバのドキュメント
// @host localhost:1323
// @BasePath /
func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out},"x-app-userid":${header:x-app-userid}}` + "\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: buildFrontendFS(),
		HTML5:      true,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := NewAppContext(c)
			return next(cc)
		}
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
			c.Response().Header().Set("Connection", "keep-alive")
			return next(c)
		}
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := UseAppContext(c)
			userId := ac.GetUserIdHeaderValue()

			if userId != "" && !ac.ExistsUser(userId) {
				return echo.ErrUnauthorized
			}

			c.Logger().Info("x-app-userid:%s", userId)

			return next(ac)
		}
	})

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	api := e.Group("/api")
	RegisterAPIHandlers(api)

	e.Logger.Fatal(e.Start(":1323"))
}
