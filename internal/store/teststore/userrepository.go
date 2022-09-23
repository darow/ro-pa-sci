package teststore

type UserRepository struct {
	store        *Store
	users        map[int]*model.User
	authAttempts map[int]*model.AuthenticationLog
}
