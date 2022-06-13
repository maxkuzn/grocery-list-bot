package model

type ItemID uint64

type Item struct {
	ID          ItemID
	Description string
	Checked     bool
}
