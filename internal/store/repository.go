package store

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}

type SessionRepository interface {
	Create(*model.Session) error
	Find(string) (*model.Session, error)
}
