package metainfo

import "github.com/maxkuzn/grocery-list-bot/internal/model"

type UserMetaInfo struct {
	AnyListSelected bool
	SelectedList    model.ListID // optional

	// TODO
	// UnfinishedCommand string // optional
}

type MetaInfoStorer struct {
	users map[model.UserID]UserMetaInfo
}

func New() *MetaInfoStorer {
	return &MetaInfoStorer{
		users: make(map[model.UserID]UserMetaInfo),
	}
}

func (s *MetaInfoStorer) Get(userID model.UserID) UserMetaInfo {
	return s.users[userID]
}

func (s *MetaInfoStorer) SelectList(userID model.UserID, listID model.ListID) {
	s.users[userID] = UserMetaInfo{
		AnyListSelected: true,
		SelectedList:    listID,
	}
}

func (s *MetaInfoStorer) SetList(userID model.UserID, list model.List) {
	s.users[userID] = UserMetaInfo{
		AnyListSelected: true,
		SelectedList:    list.ID,
		// TODO: set list
	}
}

func (s *MetaInfoStorer) UnsetList(userID model.UserID) {
	s.users[userID] = UserMetaInfo{
		AnyListSelected: false,
	}
}
