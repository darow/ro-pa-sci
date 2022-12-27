package teststore

import (
	"errors"

	"github.com/darow/ro-pa-sci/internal/model"
	"github.com/darow/ro-pa-sci/internal/store"
)

type InviteRepository struct {
	store   *Store
	invites map[int]*model.Invite
}

var inviteID int

func (r *InviteRepository) Create(invite *model.Invite) error {
	err := r.checkDuplicates(invite)
	if err != store.ErrRecordNotFound {
		return errors.New("приглашение уже было отправлено ранее")
	}

	inviteID++
	invite.ID = inviteID
	r.invites[invite.ID] = invite

	return nil
}

// checkDuplicates ErrRecordNotFound - нет дубликатов. nil - нашли дубликат
func (r *InviteRepository) checkDuplicates(invite *model.Invite) error {
	for _, inv := range r.invites {
		if inv.From == invite.From && inv.To == invite.To {
			return nil
		}
	}

	return store.ErrRecordNotFound
}

func (r *InviteRepository) Find(id int) (*model.Invite, error) {
	i, ok := r.invites[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return i, nil
}
