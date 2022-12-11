package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mktbsh/web-speed-hackathon-2022/domain/model"
	"github.com/mktbsh/web-speed-hackathon-2022/domain/service"
	"github.com/uptrace/bun"
)

type raceAPIHandlers struct {
	rs  service.RaceService
	bts service.BettingTicketService
}

type RacesParams struct {
	Since *int64 `query:"since"`
	Until *int64 `query:"until"`
}

type RacesResponse struct {
	Races []model.Race `json:"races"`
}

func unix2time(t int64) time.Time {
	if t == 0 {
		return time.Time{}
	}

	// 秒単位のUNIX時間がt, ナノ秒が0の時刻を持つtime.Time型を返す
	return time.Unix(t, 0)
}

func RegisterRaceAPI(api *echo.Group, db *bun.DB) {
	handlers := raceAPIHandlers{
		rs:  service.NewRaceService(db),
		bts: service.NewBettingTicketService(db),
	}

	api.GET("", handlers.FindRaces)
	api.GET("/:raceId", handlers.FindRaceById)

	bt := api.Group("/:raceId/betting-tickets")
	bt.GET("", handlers.GetBettingTickets)
	bt.POST("", handlers.PostBettingTicket)
}

func (h *raceAPIHandlers) FindRaces(c echo.Context) error {
	var params RacesParams
	err := c.Bind(&params)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	since := unix2time(*params.Since)
	until := unix2time(*params.Until)
	races, err := h.rs.Find(&since, &until)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, RacesResponse{
		Races: *races,
	})

}

func (h *raceAPIHandlers) FindRaceById(c echo.Context) error {
	raceId := c.Param("raceId")
	if raceId == "" {
		return echo.ErrBadRequest
	}

	race, err := h.rs.FindById(raceId)

	if err != nil || race.IsEmpty() {
		return echo.ErrNotFound
	}

	return c.JSON(200, race)
}

type BettingTicketsResponse struct {
	BettingTickets *[]model.BettingTicket `json:"bettingTickets"`
}

func (h *raceAPIHandlers) GetBettingTickets(c echo.Context) error {
	raceId := c.Param("raceId")
	if raceId == "" {
		return echo.ErrBadRequest
	}

	ac := UseAppContext(c)
	user := ac.UseState().User

	if user == nil {
		return echo.ErrUnauthorized
	}

	response := BettingTicketsResponse{
		BettingTickets: &[]model.BettingTicket{},
	}

	tickets, err := h.bts.FindByRaceIdAndUserId(raceId, user.ID)

	if err != nil {
		return c.JSON(http.StatusOK, response)
	}

	response.BettingTickets = tickets
	return c.JSON(http.StatusOK, response)

}

type PostBettingTicketBody struct {
	Type string  `json:"type"`
	Key  []int64 `json:"key"`
}

func (h *raceAPIHandlers) PostBettingTicket(c echo.Context) error {
	raceId := c.Param("raceId")
	if raceId == "" {
		return echo.ErrBadRequest
	}

	ac := UseAppContext(c)
	user := ac.UseState().User
	if user == nil {
		return echo.ErrUnauthorized
	}

	if user.Balance < 100 {
		return c.String(428, "Precondition Required")
	}

	var body PostBettingTicketBody
	err := ac.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	ticket, err := h.bts.Bet(user, raceId, body.Type, body.Key)

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, ticket)
}
