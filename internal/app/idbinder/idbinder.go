package idbinder

import (
	"errors"
	"fmt"

	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

const sep = "#"

func buildKey(userName string, listName string) string {
	return userName + "#" + listName
}

type IDBinder struct {
	tg2User map[int]model.UserID
	user2Tg map[model.UserID]int

	listName2ID map[string]model.ListID
}

func New() *IDBinder {
	return &IDBinder{
		tg2User: make(map[int]model.UserID),
		user2Tg: make(map[model.UserID]int),

		listName2ID: make(map[string]model.ListID),
	}
}

func (b *IDBinder) BindUser(tgID int, userID model.UserID) error {
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

func (b *IDBinder) Tg2User(tgID int) (userID model.UserID, ok bool) {
	userID, ok = b.tg2User[tgID]
	return
}

func (b *IDBinder) User2Tg(userID model.UserID) (tgID int, ok bool) {
	tgID, ok = b.user2Tg[userID]
	return
}

func (b *IDBinder) BindList(userName string, listName string, listID model.ListID) error {
	key := buildKey(userName, listName)
	if _, ok := b.listName2ID[key]; ok {
		return errors.New("List with this name already binded")
	}
	b.listName2ID[key] = listID
	return nil
}

func (b *IDBinder) ListName2ID(userName string, listName string) (listID model.ListID, ok bool) {
	key := buildKey(userName, listName)
	listID, ok = b.listName2ID[key]
	return
}
