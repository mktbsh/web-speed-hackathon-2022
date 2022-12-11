package model

import (
	"time"

	"github.com/uptrace/bun"
)

type BettingTicket struct {
	bun.BaseModel `bun:"table:betting_ticket,alias:bt"`

	ID        string    `bun:"id,pk,default:gen_random_uuid()"`
	Key       []int64   `bun:"key,notnull,type:json"`
	Type      string    `bun:"type,notnull,type:varchar"`
	CreatedAt time.Time `bun:"createdAt,nullzero,notnull,default:current_timestamp"`

	RaceId string `bun:"raceId,notnull,type:varcher"`
	UserId string `bun:"userId,notnull,type:varchar"`
}
