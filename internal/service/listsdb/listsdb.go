package listsdb

import "github.com/maxkuzn/grocery-list-bot/internal/model"

type ListsDB interface {
	// Lists creation
	CreateList(owner model.UserID) (model.ListID, error)
	DeleteList(listID model.ListID) error

	// List modification
	AddItem(listID model.ListID, item model.Item) error
	GetItems(listID model.ListID) ([]model.Item, error)

	GetLists(userID model.UserID) ([]model.ListID, error)
	GetOwnedLists(userID model.UserID) ([]model.ListID, error)
	GetWritableLists(userID model.UserID) ([]model.ListID, error)
	GetReadableLists(userID model.UserID) ([]model.ListID, error)

	MakeOwner(userID model.UserID, listID model.ListID) error
	AddReadAccessRights(userID model.UserID, listID model.ListID) error
	AddWriteAccessRights(userID model.UserID, listID model.ListID) error
}
