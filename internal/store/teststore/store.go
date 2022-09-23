package teststore

import "rock-paper-scissors/internal/store"

type Store struct {
	userRepository    *UserRepository
	sessionRepository *SessionRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
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
