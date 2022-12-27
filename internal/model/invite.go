package model

import "time"

type Invite struct {
	ID        int `json:"ID"`
	From      int `json:"from"`
	To        int `json:"to"`
	Timestamp time.Time
}
