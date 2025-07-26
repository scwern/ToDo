package service

import (
	"ToDo/internal/mocks"
	"github.com/google/uuid"
	"testing"

	"ToDo/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

func createTestUser() user.User {
	u := user.NewUser("Test User", "test@example.com", "password")
	u.SetID(uuid.New())
	return u
}

func TestUserServiceGetByEmail(t *testing.T) {
	repo := mocks.NewUserRepositoryInterface(t)
	service := NewUserService(repo)

	email := "test@example.com"
	expectedUser := createTestUser()

	repo.EXPECT().GetByEmail(email).Return(&expectedUser, nil).Once()

	result, err := service.GetByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, &expectedUser, result)
}

func TestUserServiceCreate(t *testing.T) {
	repo := mocks.NewUserRepositoryInterface(t)
	service := NewUserService(repo)

	newUser := user.NewUser("Test", "test@example.com", "pass")
	createdUser := newUser

	repo.EXPECT().Create(newUser).Return(createdUser, nil).Once()

	result, err := service.Create(newUser)

	assert.NoError(t, err)
	assert.Equal(t, createdUser, result)
}

func TestUserServiceGetAll(t *testing.T) {
	repo := mocks.NewUserRepositoryInterface(t)
	service := NewUserService(repo)

	expectedUsers := []user.User{createTestUser()}

	repo.EXPECT().GetAll().Return(expectedUsers, nil).Once()

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
}

func TestUserServiceGetById(t *testing.T) {
	repo := mocks.NewUserRepositoryInterface(t)
	service := NewUserService(repo)

	userID := uuid.New()
	expectedUser := createTestUser()

	repo.EXPECT().GetById(userID).Return(&expectedUser, nil).Once()

	result, err := service.GetById(userID)

	assert.NoError(t, err)
	assert.Equal(t, &expectedUser, result)
}

func TestUserServiceUpdate(t *testing.T) {
	repo := mocks.NewUserRepositoryInterface(t)
	service := NewUserService(repo)

	userID := uuid.New()
	updatedUser := createTestUser()
	updatedUser.SetName("Updated Name")

	repo.EXPECT().Update(userID, updatedUser).Return(&updatedUser, nil).Once()

	result, err := service.Update(userID, updatedUser)

	assert.NoError(t, err)
	assert.Equal(t, &updatedUser, result)
}

func TestUserServiceDelete(t *testing.T) {
	repo := mocks.NewUserRepositoryInterface(t)
	service := NewUserService(repo)

	userID := uuid.New()

	repo.EXPECT().Delete(userID).Return(nil).Once()

	err := service.Delete(userID)

	assert.NoError(t, err)
}
