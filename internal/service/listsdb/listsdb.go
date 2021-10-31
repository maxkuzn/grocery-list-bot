package listsdb

import (
	"errors"

	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

var (
	UserDoesntExist = errors.New("User doesnt exists")
	ListDoesntExist = errors.New("List doesnt exists")

	ListNameDuplicate     = errors.New("List with this name already exists")
	NotEnoughAccessRights = errors.New("Not enough access rights")
)

type ListsDB interface {
	// Above user methods
	CreateUser() model.UserID

	// List handle methods
	CreateList(userID model.UserID, name string) (model.ListID, error)
	RemoveList(userID model.UserID, listID model.ListID) error

	// Get info about users lists
	GetAllLists(userID model.UserID) ([]model.List, error)
	GetList(listID model.ListID) (model.List, error)

	// List modification methods
	AddItem(userID model.UserID, listID model.ListID, item model.Item) error
	RemoveItem(userID model.UserID, listID model.ListID, item model.ItemID) error
	ModifyItem(userID model.UserID, listID model.ListID, item model.Item) error
	// CheckItem(userID model.UserID, listID model.ListID, item model.ItemID)
	// UnckechItem(userID model.UserID, listID model.ListID, item model.ItemID)

	// TODO: Share feature
	// AddUser(ownerID model.UserID, listID model.ListID, shareWith model.UserID, readOnly bool)
	// RemoveUser(ownerID model.UserID, listID model.ListID, shareWith model.UserID)
}
