package model

import "errors"

var (
	ItemNotFoundErr = errors.New("item not found")
)

type ItemID uint64

type Item struct {
	ID          ItemID
	Checked     bool
	Description string
}

type ListID uint64

type List struct {
	ID    ListID
	Owner UserID
	Name  string
	// TODO: add share feature
	// Users   map[UserID]*User

	nextItemID ItemID
	Items      map[ItemID]Item
}

func NewList(id ListID, owner UserID, name string) *List {
	return &List{
		ID:    id,
		Owner: owner,
		Name:  name,

		nextItemID: 0,
		Items:      make(map[ItemID]Item),
	}
}

func (l *List) AddItem(item Item) {
	item.ID = l.nextItemID
	l.nextItemID++

	l.Items[item.ID] = item
}

func (l *List) RemoveItem(itemID ItemID) error {
	_, ok := l.Items[itemID]
	if !ok {
		return ItemNotFoundErr
	}
	delete(l.Items, itemID)
	return nil
}

func (l *List) ModifyItem(item Item) error {
	_, ok := l.Items[item.ID]
	if !ok {
		return ItemNotFoundErr
	}
	l.Items[item.ID] = item
	return nil
}
