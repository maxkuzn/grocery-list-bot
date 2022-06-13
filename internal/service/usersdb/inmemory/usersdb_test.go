package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maxkuzn/grocery-list-bot/internal/model"
	"github.com/maxkuzn/grocery-list-bot/internal/service/usersdb"
)

func TestInMemoryUsersDB(t *testing.T) {
	t.Parallel()

	type putStruct struct {
		value   int64
		want    model.UserID
		wantErr error
	}

	type getStruct struct {
		value   int64
		want    model.UserID
		wantErr error
	}

	testCases := []struct {
		name string
		put  []putStruct
		get  []getStruct
	}{
		{
			name: "simple",
			put: []putStruct{
				{value: 1, want: 1},
			},
			get: []getStruct{
				{value: 1, want: 1},
			},
		},
		{
			name: "several users",
			put: []putStruct{
				{value: 1, want: 1},
				{value: 5, want: 2},
				{value: 100, want: 3},
				{value: 42, want: 4},
			},
			get: []getStruct{
				{value: 42, want: 4},
				{value: 1, want: 1},
				{value: 100, want: 3},
				{value: 100, want: 3},
				{value: 42, want: 4},
				{value: 42, want: 4},
				{value: 5, want: 2},
				{value: 1, want: 1},
			},
		},
		{
			name: "second put",
			put: []putStruct{
				{value: 1, want: 1},
				{value: 1, wantErr: usersdb.ErrUserAlreadyExists},
			},
			get: []getStruct{
				{value: 1, want: 1},
			},
		},
		{
			name: "unknown user",
			put: []putStruct{
				{value: 1, want: 1},
			},
			get: []getStruct{
				{value: 1, want: 1},
				{value: 2, wantErr: usersdb.ErrNotFound},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			db := NewInMemoryUsersDB()

			for _, put := range tc.put {
				got, err := db.CreateUser(put.value)
				if put.wantErr != nil {
					assert.ErrorIs(t, err, put.wantErr)
				} else {
					assert.Equal(t, put.want, got)
				}
			}

			for _, get := range tc.get {
				got, err := db.GetUserID(get.value)
				if get.wantErr != nil {
					assert.ErrorIs(t, err, get.wantErr)
				} else {
					assert.Equal(t, get.want, got)
				}
			}
		})
	}
}
