package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mktbsh/web-speed-hackathon-2022/domain/model"
	"github.com/uptrace/bun"
)

type BettingTicketService struct {
	db *bun.DB
}

func NewBettingTicketService(db *bun.DB) BettingTicketService {
	return BettingTicketService{db}
}

func (s *BettingTicketService) FindByRaceIdAndUserId(raceId string, userId string) (*[]model.BettingTicket, error) {
	tickets := []model.BettingTicket{}

	err := s.db.NewSelect().Model(&tickets).Where("raceId = ?", raceId).Where("userId = ?", userId).Scan(context.Background())

	return &tickets, err
}

func (s *BettingTicketService) Bet(user *model.User, raceId string, betType string, key []int64) (*model.BettingTicket, error) {
	ticket := model.BettingTicket{
		ID:     uuid.NewString(),
		Key:    key,
		Type:   betType,
		RaceId: raceId,
		UserId: user.ID,
	}

	_, err := s.db.NewInsert().Model(&ticket).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	userService := NewUserService(s.db)
	user.Charge(-100)
	go userService.Save(user)

	return &ticket, nil
}
