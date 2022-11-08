package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Race struct {
	bun.BaseModel `bun:"table:race,alias:r"`

	ID      string    `bun:"id,pk,default:gen_random_uuid()" json:"id"`
	CloseAt time.Time `bun:"closeAt,notnull,type:datetime" json:"closeAt"`
	Image   string    `bun:"image,notnull,type:varchar" json:"image"`
	Name    string    `bun:"name,notnull,type:varchar" json:"name"`
	StartAt time.Time `bun:"startAt,notnull,type:datetime" json:"startAt"`

	Entries      []*RaceEntry `bun:"rel:has-many,join:id=raceId" json:"entries"`
	TrifectaOdds []*OddsItem  `bun:"rel:has-many,join:id=raceId" json:"trifectaOdds"`
}

func (r *Race) IsEmpty() bool {
	return r.ID == ""
}
