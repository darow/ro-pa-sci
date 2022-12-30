package store

import "github.com/darow/ro-pa-sci/internal/model"

type UserRepository interface {
	Create(*model.User) error
	Find(id int) (*model.User, error)
	Login(login, password string) (*model.User, error)
	GetTop() []*model.User
}

type SessionRepository interface {
	Create(*model.User) (*model.Session, error)
	Find(string) (*model.Session, error)
}

type PlayerRepository interface {
	New(string) (*model.Session, error)
	All() (*model.Session, error)
}

type InviteRepository interface {
	Create(*model.Invite) error
	Update(invite *model.Invite) error
	Get(id int) (*model.Invite, error)
	GetByUser(int) ([]*model.Invite, error)
}

type GameRepository interface {
	Create(game *model.Game) error
	GetByUser(userID int) ([]*model.Game, error)
}
