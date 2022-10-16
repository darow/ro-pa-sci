package teststore

import (
	"errors"
	"github.com/darow/ro-pa-sci/internal/model"
	"github.com/darow/ro-pa-sci/internal/store"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	store       *Store
	users       map[int]*model.User
	onlineUsers []*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	_, err := r.findByLogin(u.Username)
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
	if err = bcrypt.CompareHashAndPassword(u.EncryptedPassword, []byte(password)); err != nil {
		return nil, errors.New("Пароль не подходит")
	}

	return u, nil
}

func (r *UserRepository) findByLogin(login string) (*model.User, error) {
	for _, u := range r.users {
		if u.Username == login {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) PopOnlineUser(u *model.User) {
	for i := range r.onlineUsers {
		if r.onlineUsers[i].ID == u.ID {
			r.onlineUsers[i], r.onlineUsers[len(r.onlineUsers)-1] = r.onlineUsers[len(r.onlineUsers)-1], r.onlineUsers[i]
		}
	}
}

func (r *UserRepository) GetTop() []*model.User {
	res := make([]*model.User, 0, len(r.users))
	for _, v := range r.users {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Score > res[j].Score
	})

	return res
}
