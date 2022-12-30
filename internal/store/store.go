package store

type Store interface {
	User() UserRepository
	Session() SessionRepository
	Invite() InviteRepository
	Game() GameRepository
}
