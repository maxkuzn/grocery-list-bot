package model

type ItemID uint64

type Item struct {
	ID          ItemID
	Description string
}

type ListID uint64

type List struct {
	ID    ListID
	Owner UserID
	Name  string
	// TODO: add share feature
	// Users   map[UserID]*User

	Items []Item
}

func NewList(id ListID, owner UserID, name string) *List {
	return &List{
		ID:    id,
		Owner: owner,
		Name:  name,
	}
}
