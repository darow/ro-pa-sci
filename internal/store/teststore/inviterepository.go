package teststore

import (
	"errors"
	"time"

	"github.com/darow/ro-pa-sci/internal/model"
	"github.com/darow/ro-pa-sci/internal/store"
)

type InviteRepository struct {
	store   *Store
	invites map[int]*model.Invite
}

var inviteIDIncrement int

func (r *InviteRepository) Create(invite *model.Invite) error {
	err := r.checkDuplicates(invite)
	if err != store.ErrRecordNotFound {
		return errors.New("приглашение уже было отправлено ранее")
	}

	invite.Created = time.Now()
	invite.Decision = model.DecisionNotDecided
	inviteIDIncrement++
	invite.ID = inviteIDIncrement
	r.invites[invite.ID] = invite

	return nil
}

func (r *InviteRepository) Update(invite *model.Invite) error {
	r.invites[invite.ID] = invite

	return nil
}

// checkDuplicates ErrRecordNotFound - нет дубликатов. nil - нашли дубликат
func (r *InviteRepository) checkDuplicates(invite *model.Invite) error {
	for _, inv := range r.invites {
		if inv.Decision == model.DecisionNotDecided && inv.From == invite.From && inv.To == invite.To {
			return nil
		}
	}

	return store.ErrRecordNotFound
}

func (r *InviteRepository) GetByUser(userID int) ([]*model.Invite, error) {
	res := make([]*model.Invite, 0, 1)
	for _, inv := range r.invites {
		if userID == inv.To || userID == inv.From {
			res = append(res, inv)
		}
	}

	return res, nil
}

//
//func (r *InviteRepository) Decide(inviteID int, decision uint8) error {
//	inv, err := r.Get(inviteID)
//	if err != nil {
//		return err
//	}
//
//	inv.Decision = decision
//	return nil
//}

func (r *InviteRepository) Get(id int) (*model.Invite, error) {
	i, ok := r.invites[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return i, nil
}
