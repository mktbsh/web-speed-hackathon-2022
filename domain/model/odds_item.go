package model

import (
	"encoding/json"
	"strconv"

	"github.com/uptrace/bun"
)

type OddsItem struct {
	bun.BaseModel `bun:"table:odds_item,alias:oi"`

	ID   string  `bun:"id,pk,default:gen_random_uuid()" json:"id"`
	Key  []int64 `bun:"key,notnull,type:json" json:"key"`
	Odds string  `bun:"odds,notnull,type:int" json:"odds"`
	Type string  `bun:"type,notnull,type:varchar" json:"type"`

	RaceId string `bun:"raceId,notnull,type:varcher" json:"raceId"`
}

func (o *OddsItem) MarshalJSON() ([]byte, error) {
	odds, err := strconv.ParseFloat(o.Odds, 64)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&struct {
		ID     string  `json:"id"`
		Key    []int64 `json:"key"`
		Odds   float64 `json:"odds"`
		Type   string  `json:"type"`
		RaceId string  `json:"raceId"`
	}{
		ID:     o.ID,
		Key:    o.Key,
		Odds:   odds,
		Type:   o.Type,
		RaceId: o.RaceId,
	})
}
