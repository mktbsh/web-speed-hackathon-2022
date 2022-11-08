package model

import "github.com/uptrace/bun"

type Player struct {
	bun.BaseModel `bun:"table:player,alias:p"`

	ID        string `bun:"id,pk,default:gen_random_uuid()"`
	Image     string `bun:"image,notnull,type:text"`
	Name      string `bun:"name,notnull,type:varchar"`
	ShortName string `bun:"shortName,notnull,type:varchar"`
}
