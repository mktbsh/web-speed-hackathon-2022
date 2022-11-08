package main

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mktbsh/web-speed-hackathon-2022/database"
)

type HeroImageResponse struct {
	Hash string `json:"hash"`
	Url  string `json:"url"`
}

func RegisterAPIHandlers(api *echo.Group) {

	api.POST("/initialize", func(c echo.Context) error {
		return database.InitializeDatabase(true)
	})

	api.GET("/hero", func(c echo.Context) error {
		url := "/assets/images/hero.jpg"
		md5 := md5.Sum([]byte(url))

		return c.JSON(http.StatusOK, HeroImageResponse{
			Hash: fmt.Sprintf("%x", md5),
			Url:  url,
		})
	})

	users := api.Group("/users")
	RegisterUserAPI(users, database.DB)

	races := api.Group("/races")
	RegisterRaceAPI(races, database.DB)
}
