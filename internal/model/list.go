package model

type ListID uint64

type List struct {
	ID    ListID
	Owner UserID
	Name  string
}
