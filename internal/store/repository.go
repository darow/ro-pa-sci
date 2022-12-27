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
}
