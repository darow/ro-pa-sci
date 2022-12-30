package model

import "time"

type Game struct {
	ID      int
	Players [2]int
	Winner  int
	Created time.Time
}
