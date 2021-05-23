package mock

import (
	"bruqus/snippetbox/pkg/models"
	"time"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Keks",
	Email:   "test@example.ru",
	Created: time.Now(),
	Active:  true,
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "test@example.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
