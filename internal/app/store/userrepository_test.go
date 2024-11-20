package store_test

import (
	"github.com/amangeldi0/http-rest-api/internal/app/model"
	"github.com/amangeldi0/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)

	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "test@test.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)

	defer teardown("users")

	email := "test@test.com"

	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "test@test.com",
	})

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)

}
