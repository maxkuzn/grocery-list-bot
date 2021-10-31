package model

import (
	"errors"
	"fmt"
)

type UserID uint64

type AccessRights uint8

const (
	Owner AccessRights = iota
	Write
	Read
)

func (r AccessRights) String() string {
	switch r {
	case Owner:
		return "owner access rights"
	case Write:
		return "write access rights"
	case Read:
		return "read access rights"
	default:
		panic("Unknown access rights")
	}
}

var (
	ListNameDuplicate     = errors.New("List with this name already exists")
	NotEnoughAccessRights = errors.New("Not enough access rights")
)

type User struct {
	ID          UserID
	listsRights map[ListID]AccessRights
	ownedLists  map[string]ListID
}

func NewUser(id UserID) *User {
	return &User{
		ID:          id,
		listsRights: make(map[ListID]AccessRights),
		ownedLists:  make(map[string]ListID),
	}
}

func (u *User) MakeOwner(list *List) error {
	rights, ok := u.listsRights[list.ID]
	if ok {
		return fmt.Errorf("User [id=%d] already have list [id=%d] with rights %v",
			u.ID, list.ID, rights)
	}
	_, ok = u.ownedLists[list.Name]
	if ok {
		return ListNameDuplicate
	}
	u.listsRights[list.ID] = Owner
	u.ownedLists[list.Name] = list.ID
	return nil
}

func (u *User) Remove(list *List) error {
	rights, ok := u.listsRights[list.ID]
	if !ok {
		panic("Remove of unexisting list")
	}
	if rights != Owner {
		return NotEnoughAccessRights
	}
	delete(u.listsRights, list.ID)
	delete(u.ownedLists, list.Name)
	return nil
}

func (u *User) ListIDByName(name string) (listID ListID, ok bool) {
	listID, ok = u.ownedLists[name]
	return
}
