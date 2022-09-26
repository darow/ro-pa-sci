package store

import "rock-paper-scissors/internal/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}

type SessionRepository interface {
	Create(*model.User) error
	Find(string) (*model.Session, error)
}
