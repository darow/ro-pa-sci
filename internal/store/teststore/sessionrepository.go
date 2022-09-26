package teststore

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"rock-paper-scissors/internal/model"
	"rock-paper-scissors/internal/store"
)

type SessionRepository struct {
	store    *Store
	sessions map[int]*model.Session
}

const (
	sessionLiveTimeShort = time.Second * time.Duration(20)
	sessionLiveTime      = time.Minute * time.Duration(30)
	sessionLiveTimeLong  = time.Hour * time.Duration(1000)
)

// CreateToken вспомогательная функция создания токена для сессии.
func (r *SessionRepository) createToken() string {
	b := md5.Sum([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	token := hex.EncodeToString(b[:])

	// Проверяем, что md5 сгенерировал токен, которого у нас еще нет.
	_, err := r.Find(token)
	if err == store.ErrRecordNotFound {
		return token
	}

	return r.createToken()
}

func (r *SessionRepository) Create(u *model.User) error {
	s := &model.Session{
		UserID:         u.ID,
		Token:          r.createToken(),
		ExpirationTime: time.Now().Local().Add(sessionLiveTimeLong),
	}
	s.ID = len(r.sessions)
	r.sessions[s.ID] = s
	return nil
}

func (r *SessionRepository) Find(token string) (*model.Session, error) {
	for _, s := range r.sessions {
		if s.Token == token {
			return s, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
