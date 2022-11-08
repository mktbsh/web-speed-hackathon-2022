package service

import (
	"context"

	"github.com/mktbsh/web-speed-hackathon-2022/domain/model"
	"github.com/uptrace/bun"
)

type UserService struct {
	db *bun.DB
}

func NewUserService(db *bun.DB) UserService {
	return UserService{db}
}

func (s *UserService) Save(user *model.User) (*model.User, error) {
	_, err := s.db.NewInsert().Model(user).
		On("CONFLICT (id) DO UPDATE").
		Set("balance = EXCLUDED.balance").
		Set("payoff = EXCLUDED.payoff").
		Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}
