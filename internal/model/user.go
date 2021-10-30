package model

type UserID uint64

type User struct {
	ID    UserID
	Lists []*List
}

func NewUser(id UserID) *User {
	return &User{
		ID: id,
	}
}
