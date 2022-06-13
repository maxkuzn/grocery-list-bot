package inmemory

import (
	"sync"

	"github.com/maxkuzn/grocery-list-bot/internal/model"
	"github.com/maxkuzn/grocery-list-bot/internal/service/usersdb"
)

type inMemoryUsersDB struct {
	nextUserID model.UserID

	m     sync.RWMutex
	tg2id map[int64]model.UserID
}

func NewInMemoryUsersDB() *inMemoryUsersDB {
	return &inMemoryUsersDB{
		nextUserID: 1,
		tg2id:      make(map[int64]model.UserID),
	}
}

func (db *inMemoryUsersDB) CreateUser(telegramID int64) (model.UserID, error) {
	db.m.Lock()
	defer db.m.Unlock()

	_, ok := db.tg2id[telegramID]
	if ok {
		return 0, usersdb.ErrUserAlreadyExists
	}

	userID := db.nextUserID
	db.nextUserID++

	db.tg2id[telegramID] = userID

	return userID, nil
}

func (db *inMemoryUsersDB) GetUserID(telegramID int64) (model.UserID, error) {
	db.m.RLock()
	defer db.m.RUnlock()

	userID, ok := db.tg2id[telegramID]
	if !ok {
		return 0, usersdb.ErrNotFound
	}

	return userID, nil
}
