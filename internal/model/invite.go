package model

import "time"

const (
	DecisionDeclined = iota
	DecisionAccepted
	DecisionNotDecided
)

var Decisions = map[uint8]struct{}{
	DecisionDeclined:   struct{}{},
	DecisionAccepted:   struct{}{},
	DecisionNotDecided: struct{}{},
}

type Invite struct {
	ID       int       `json:"id"`
	From     int       `json:"from"`
	To       int       `json:"to"`
	Created  time.Time `json:"created"`
	Decision uint8     `json:"decision"`
}
