package usersdb

import (
	"errors"

	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

var (
	ErrNotFound          = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UsersDB interface {
	CreateUser(telegramID int64) (model.UserID, error)
	GetUserID(telegramID int64) (model.UserID, error)
}
