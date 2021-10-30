package listsdb

import "github.com/maxkuzn/grocery-list-bot/internal/model"

type ListsDB interface {
	// Above user methods
	CreateUser() model.UserID

	// List handle methods
	CreateList(userID model.UserID) model.ListID
	RemoveList(userID model.UserID, listID model.ListID)

	// Get info about users lists
	GetAllLists(userID model.UserID) []*model.List
	GetList(listID model.ListID) *model.List

	// List modification methods
	AddItem(userID model.UserID, listID model.ListID, item model.Item)
	RemoveItem(userID model.UserID, listID model.ListID, item model.ItemID)
	CheckItem(userID model.UserID, listID model.ListID, item model.ItemID)
	UnckechItem(userID model.UserID, listID model.ListID, item model.ItemID)

	// TODO: Share feature
	// AddUser(ownerID model.UserID, listID model.ListID, shareWith model.UserID, readOnly bool)
	// RemoveUser(ownerID model.UserID, listID model.ListID, shareWith model.UserID)
}
