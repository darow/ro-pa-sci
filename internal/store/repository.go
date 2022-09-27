package store

import "rock-paper-scissors/internal/model"

type UserRepository interface {
	Create(*model.User) error
	Find(id int) (*model.User, error)
	Login(login, password string) (*model.User, error)
}

type SessionRepository interface {
	Create(*model.User) (*model.Session, error)
	Find(string) (*model.Session, error)
}
