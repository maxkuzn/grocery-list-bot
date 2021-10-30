package model

type ItemID uint64

type Item struct {
	ID          ItemID
	Description string
}

type ListID uint64

type List struct {
	ID      ListID
	OwnerID UserID
	Owner   *User
	// TODO: add share feature
	// Users   map[UserID]*User

	Items []*Item
}

func NewList(id ListID, owner *User) *List {
	return &List{
		ID:      id,
		OwnerID: owner.ID,
		Owner:   owner,
	}
}
