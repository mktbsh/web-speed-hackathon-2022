package model

import "github.com/uptrace/bun"

type RaceEntry struct {
	bun.BaseModel `bun:"table:race_entry,alias:re"`

	ID             string  `bun:"id,pk,default:gen_random_uuid()" json:"id"`
	Comment        string  `bun:"comment,notnull" json:"comment"`
	First          int64   `bun:"first,notnull,type:integer" json:"first"`
	FirstRate      float64 `bun:"firstRate,notnull,type:float" json:"firstRate"`
	Number         int64   `bun:"number,notnull,type:integer" json:"number"`
	Others         int64   `bun:"others,notnull,type:integer" json:"others"`
	PaperWin       int64   `bun:"paperWin,notnull,type:integer" json:"paperWin"`
	PredictionMark string  `bun:"predictionMark,notnull,type:varchar" json:"predictionMark"`
	RockWin        int64   `bun:"rockWin,notnull,type:integer" json:"rockWin"`
	ScissorsWin    int64   `bun:"scissorsWin,notnull,type:integer" json:"scissorsWin"`
	Second         int64   `bun:"second,notnull,type:integer" json:"second"`
	Third          int64   `bun:"third,notnull,type:integer" json:"third"`
	ThirdRate      float64 `bun:"thirdRate,notnull,type:float" json:"thirdRate"`

	PlayerId string  `bun:"playerId,type:varchar" json:"playerId"`
	RaceId   string  `bun:"raceId,type:varchar" json:"raceId"`
	Player   *Player `bun:"rel:belongs-to,join:playerId=id" json:"player"`
	Race     *Race   `bun:"rel:belongs-to,join:raceId=id" json:"race"`
}
