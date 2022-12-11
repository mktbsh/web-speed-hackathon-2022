package service

import (
	"context"
	"time"

	"github.com/mktbsh/web-speed-hackathon-2022/domain/model"
	"github.com/uptrace/bun"
)

type RaceService struct {
	db *bun.DB
}

func NewRaceService(db *bun.DB) RaceService {
	return RaceService{db}
}

func (s *RaceService) Find(since *time.Time, until *time.Time) (*[]model.Race, error) {
	races := []model.Race{}

	q := s.db.NewSelect().
		Model(&races).
		Relation("Entries", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Relation("Player")
		}).
		Relation("TrifectaOdds")

	q = startAtBetween(q, since, until)

	err := q.OrderExpr("startAt ASC").Scan(context.Background())

	return &races, err
}

func (s *RaceService) FindById(raceId string) (model.Race, error) {
	race := model.Race{}

	err := s.db.NewSelect().Model(&race).Relation("Entries", func(sq *bun.SelectQuery) *bun.SelectQuery {
		return sq.Relation("Player")
	}).
		Relation("TrifectaOdds").
		Where("id = ?", raceId).
		Scan(context.Background())

	return race, err

}

func startAtBetween(q *bun.SelectQuery, since *time.Time, until *time.Time) *bun.SelectQuery {
	dateFmt := "2006-01-02 15:04:05"

	hasSince := !since.IsZero()
	hasUntil := !until.IsZero()

	if hasSince && hasUntil {
		return q.Where("startAt BETWEEN ? AND ?", since.Format(dateFmt), until.Format(dateFmt))
	}

	if hasSince {
		return q.Where("startAt >= ?", since.Format(dateFmt))
	}

	if hasUntil {
		return q.Where("startAt <= ?", until.Format(dateFmt))
	}

	return q
}
