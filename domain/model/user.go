package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`

	ID      string `bun:"id,pk,default:gen_random_uuid()" json:"id"`
	Balance int64  `bun:"balance,notnull,default:0" json:"balance"`
	Payoff  int64  `bun:"payoff,notnull,default:0" json:"payoff"`
}

func NewUser() *User {
	return &User{
		ID:      uuid.NewString(),
		Balance: 0,
		Payoff:  0,
	}
}

func (u *User) Charge(amount int64) {
	u.Balance += int64(amount)
}
