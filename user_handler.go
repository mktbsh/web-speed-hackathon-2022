package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mktbsh/web-speed-hackathon-2022/domain/model"
	"github.com/mktbsh/web-speed-hackathon-2022/domain/service"
	"github.com/uptrace/bun"
)

type userAPIHandlers struct {
	service service.UserService
}

type ChargeBalanceForm struct {
	AccountNo  string `json:"accountNo"`
	Amount     int64  `json:"amount"`
	BankCode   string `json:"bankCode"`
	BranchCode string `json:"branchCode"`
}

func RegisterUserAPI(api *echo.Group, db *bun.DB) {
	handers := userAPIHandlers{
		service: service.NewUserService(db),
	}

	api.GET("/me", handers.GetUser)
	api.POST("/me/charge", handers.ChargeBalance)
}

func (h *userAPIHandlers) GetUser(c echo.Context) error {
	ac := UseAppContext(c)
	state := ac.UseState()

	if state.User != nil {
		return c.JSON(http.StatusOK, state.User)
	}

	user := model.NewUser()

	_, err := h.service.Save(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userAPIHandlers) ChargeBalance(c echo.Context) error {
	ac := UseAppContext(c)
	state := ac.UseState()
	user := state.User
	if user == nil {
		return echo.ErrUnauthorized
	}

	var body ChargeBalanceForm
	err := ac.Bind(&body)
	if err != nil || body.Amount <= 0 {
		ac.Logger().Error(err)
		return echo.ErrBadRequest
	}

	user.Charge(body.Amount)

	_, err = h.service.Save(user)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.NoContent(204)
}
