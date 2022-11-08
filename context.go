package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/mktbsh/web-speed-hackathon-2022/database"
	"github.com/mktbsh/web-speed-hackathon-2022/domain/model"
)

type AppContext struct {
	echo.Context
}

type AppState struct {
	User *model.User
}

const (
	X_APP_USERID = "x-app-userid"
)

func NewAppContext(c echo.Context) *AppContext {
	cc := &AppContext{c}
	return cc
}

func UseAppContext(c echo.Context) *AppContext {
	return c.(*AppContext)
}

func (c *AppContext) UseState() AppState {
	return AppState{
		User: getUserFromHeader(c),
	}
}

func (c *AppContext) GetUserIdHeaderValue() string {
	return c.Request().Header.Get(X_APP_USERID)
}

func (c *AppContext) ExistsUser(userId string) bool {
	exists, err := database.DB.NewSelect().Model((*model.User)(nil)).Where("id = ?", userId).Exists(context.Background())
	if err != nil {
		return false
	}

	return exists
}

// ----------------------------------------------------
// request
// ----------------------------------------------------
func getUserFromHeader(c echo.Context) *model.User {
	userId := c.Request().Header.Get(X_APP_USERID)
	if userId == "" {
		return nil
	}

	user := new(model.User)
	err := database.DB.NewSelect().Model(user).Where("id = ?", userId).Scan(context.Background())
	if err != nil {
		return nil
	}

	if user.ID == "" {
		return nil
	}

	return user
}
