package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-,omitempty"`
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

	return nil
}

func encryptString(s string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}

// Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
	u.EncryptedPassword = ""
}
