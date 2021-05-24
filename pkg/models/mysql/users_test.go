package mysql

import (
	"bruqus/snippetbox/pkg/models"
	"reflect"
	"testing"
	"time"
)

func TestUserModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name          string
		userID        int
		expectedUser  *models.User
		expectedError error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			expectedUser: &models.User{
				ID:      1,
				Name:    "Keks Baz",
				Email:   "test@example.com",
				Created: time.Date(2021, 5, 24, 13, 32, 30, 0, time.UTC),
				Active:  true,
			},
		},
		{
			name:          "Zero ID",
			userID:        0,
			expectedUser:  nil,
			expectedError: models.ErrNoRecord,
		},
		{
			name:          "Non-existent ID",
			userID:        2,
			expectedUser:  nil,
			expectedError: models.ErrNoRecord,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()
			m := UserModel{db}
			user, err := m.Get(test.userID)
			if err != test.expectedError {
				t.Errorf("expected %v; got %s", test.expectedError, err)
			}
			if !reflect.DeepEqual(user, test.expectedUser) {
				t.Errorf("expected %v; got %v", test.expectedUser, user)
			}
		})
	}
}
