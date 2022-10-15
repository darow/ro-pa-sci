package model

import (
	"time"
)

type Session struct {
	ID             int       `json:"-"`
	UserID         int       `json:"-"`
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expire_time"`
}
