package teststore

import (
	"errors"

	"rock-paper-scissors/internal/model"
	"rock-paper-scissors/internal/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	_, err := r.findByLogin(u.Login)
	if err == nil {
		return errors.New("пользователь с таким именем уже существует")
	}

	u.BeforeCreate()
	u.ID = len(r.users)
	r.users[u.ID] = u
	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

func (r *UserRepository) Login(login, password string) (*model.User, error) {
	u, err := r.findByLogin(login)
	if err != nil {
		return nil, err
	}
	if u.Password != password {
		return nil, errors.New("Пароль не подходит")
	}

	return u, nil
}

func (r *UserRepository) findByLogin(login string) (*model.User, error) {
	for _, u := range r.users {
		if u.Login == login {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
