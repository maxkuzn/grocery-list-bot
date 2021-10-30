package idbinder

import (
	"fmt"

	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

type IDBinder struct {
	tg2User map[int64]model.UserID
	user2Tg map[model.UserID]int64
}

func New() *IDBinder {
	return &IDBinder{
		tg2User: make(map[int64]model.UserID),
		user2Tg: make(map[model.UserID]int64),
	}
}

func (b *IDBinder) BindUser(tgID int64, userID model.UserID) error {
	if _, ok := b.tg2User[tgID]; ok {
		return fmt.Errorf("Tg user %d already binded", tgID)
	}
	if _, ok := b.user2Tg[userID]; ok {
		return fmt.Errorf("Model user %d already binded", userID)
	}
	b.tg2User[tgID] = userID
	b.user2Tg[userID] = tgID
	return nil
}

func (b *IDBinder) Tg2User(tgID int64) (userID model.UserID, ok bool) {
	userID, ok = b.tg2User[tgID]
	return
}

func (b *IDBinder) User2Tg(userID model.UserID) (tgID int64, ok bool) {
	tgID, ok = b.user2Tg[userID]
	return
}

func (b *IDBinder) Tg2ModelList() {
}

func (b *IDBinder) Model2TgList() {
}
