package teststore

import (
	"github.com/darow/ro-pa-sci/internal/model"
	"github.com/darow/ro-pa-sci/internal/store"
)

type Store struct {
	userRepository    *UserRepository
	sessionRepository *SessionRepository
	inviteRepository  *InviteRepository
	gameRepository    *GameRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store:       s,
		users:       make(map[int]*model.User),
		onlineUsers: []*model.User{},
	}
	return s.userRepository
}

func (s *Store) Session() store.SessionRepository {
	if s.sessionRepository != nil {
		return s.sessionRepository
	}

	s.sessionRepository = &SessionRepository{
		store:    s,
		sessions: make(map[int]*model.Session),
	}
	return s.sessionRepository
}

func (s *Store) Invite() store.InviteRepository {
	if s.inviteRepository != nil {
		return s.inviteRepository
	}

	s.inviteRepository = &InviteRepository{
		store:   s,
		invites: make(map[int]*model.Invite),
	}
	return s.inviteRepository
}

func (s *Store) Game() store.GameRepository {
	if s.gameRepository != nil {
		return s.gameRepository
	}

	s.gameRepository = &GameRepository{
		store: s,
		games: make(map[int]*model.Game),
	}
	return s.gameRepository
}
