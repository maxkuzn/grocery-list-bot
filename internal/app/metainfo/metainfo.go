package metainfo

import "github.com/maxkuzn/grocery-list-bot/internal/model"

type UserMetaInfo struct {
	User            model.UserID
	AnyListSelected bool
	SelectedList    model.ListID // optional

	UnfinishedCommand string // optional
}

type MetaInfoStorer struct {
	users map[model.UserID]*UserMetaInfo
}

func New() *MetaInfoStorer {
	return &MetaInfoStorer{
		users: make(map[model.UserID]*UserMetaInfo),
	}
}

func (s *MetaInfoStorer) GetLastAction(userID model.UserID) *UserMetaInfo {
	return s.users[userID]
}

func (s *MetaInfoStorer) SetLastAction(userID model.UserID, info *UserMetaInfo) {
	s.users[userID] = info
}
