package teststore

import (
	"time"

	"github.com/darow/ro-pa-sci/internal/model"
)

type GameRepository struct {
	store *Store
	games map[int]*model.Game
}

var gameIDIncrement int

func (r *GameRepository) Create(game *model.Game) error {
	game.Created = time.Now()
	gameIDIncrement++
	game.ID = inviteIDIncrement
	r.games[game.ID] = game

	return nil
}

func (r *GameRepository) GetByUser(userID int) ([]*model.Game, error) {
	res := make([]*model.Game, 0, 1)
	for _, g := range r.games {
		if userID == g.Players[0] || userID == g.Players[1] {
			res = append(res, g)
		}
	}

	return res, nil
}
