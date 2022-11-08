package domain

import "github.com/uptrace/bun"

type Client struct {
	bun.DB
}

func NewClient(db bun.DB) *Client {
	return &Client{db}
}
