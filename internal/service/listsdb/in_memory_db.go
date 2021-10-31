package listsdb

import (
	"errors"
	"fmt"
	"log"

	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

var (
	_ ListsDB = (*InMemoryDB)(nil)
)

type InMemoryDB struct {
	nextUserID model.UserID
	nextListID model.ListID

	users map[model.UserID]*model.User
	lists map[model.ListID]*model.List
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		nextUserID: 0,
		nextListID: 0,

		users: make(map[model.UserID]*model.User),
		lists: make(map[model.ListID]*model.List),
	}
}

// General
func (db *InMemoryDB) CreateUser() model.UserID {
	userID := db.nextUserID
	db.nextUserID++
	db.users[userID] = model.NewUser(userID)
	return userID
}

// Lists
func (db *InMemoryDB) CreateList(userID model.UserID, name string) (listID model.ListID, err error) {
	user, ok := db.users[userID]
	if !ok {
		err = fmt.Errorf("User with id %d doesn't exist", userID)
		return
	}

	listID = db.nextListID
	defer func() {
		if err == nil {
			db.nextListID++
		}
	}()

	list := model.NewList(listID, userID, name)
	err = user.MakeOwner(list)
	if err != nil {
		if errors.Is(err, model.ListNameDuplicate) {
			err = ListNameDuplicate
		}
		return
	}
	db.lists[listID] = list
	log.Printf("Created list with id=%d owner_id=%d", listID, userID)
	return
}

func (db *InMemoryDB) RemoveList(userID model.UserID, listID model.ListID) error {
	list, ok := db.lists[listID]
	if !ok {
		return ListDoesntExist
	}
	if list.Owner != userID {
		return NotEnoughAccessRights
	}
	user := db.users[userID]
	user.Remove(list)
	delete(db.lists, listID)
	return nil
}

// Info about lists
func (db *InMemoryDB) GetAllLists(userID model.UserID) ([]model.List, error) {
	panic("Not implemented")
}

func (db *InMemoryDB) GetList(userID model.UserID, listID model.ListID) (list model.List, err error) {
	listPtr, ok := db.lists[listID]
	if !ok {
		err = ListDoesntExist
		return
	}
	// TODO: check access rights
	if list.Owner != userID {
		err = NotEnoughAccessRights
		return
	}
	list = *listPtr
	return
}

// List modification
func (db *InMemoryDB) AddItem(userID model.UserID, listID model.ListID, item model.Item) error {
	list, ok := db.lists[listID]
	if !ok {
		return ListDoesntExist
	}
	// TODO: check access rights
	if list.Owner != userID {
		return NotEnoughAccessRights
	}

	list.AddItem(item)
	return nil
}

func (db *InMemoryDB) RemoveItem(userID model.UserID, listID model.ListID, itemID model.ItemID) error {
	list, ok := db.lists[listID]
	if !ok {
		return ListDoesntExist
	}
	// TODO: check access rights
	if list.Owner != userID {
		return NotEnoughAccessRights
	}

	return list.RemoveItem(itemID)
}

func (db *InMemoryDB) ModifyItem(userID model.UserID, listID model.ListID, item model.Item) error {
	list, ok := db.lists[listID]
	if !ok {
		return ListDoesntExist
	}
	// TODO: check access rights
	if list.Owner != userID {
		return NotEnoughAccessRights
	}

	return list.ModifyItem(item)
}
