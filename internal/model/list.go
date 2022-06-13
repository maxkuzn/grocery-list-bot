package model

import (
	"errors"
)

var (
	ItemNotFoundErr = errors.New("item not found")
)

type ListID uint64

type List struct {
	ID    ListID
	Owner UserID
	Name  string
}
