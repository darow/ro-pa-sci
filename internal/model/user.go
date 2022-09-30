package model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword []byte `json:"-"`
	IsOnline          bool   `json:"is_online"`
	Score             int    `json:"score"`
}

// BeforeCreate TODO Сюда можно добавить валидацию
func (u *User) BeforeCreate() error {
	if len(u.Password) == 0 {
		return errors.New("пароль не может быть пустым")
	}

	enc, err := encryptString(u.Password)
	if err != nil {
		return err
	}

	u.EncryptedPassword = enc
	u.Password = ""
	u.Score = int(time.Now().UnixMicro() % 10)

	return nil
}

func encryptString(s string) ([]byte, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	return enc, nil
}

// Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
	u.EncryptedPassword = nil
}
