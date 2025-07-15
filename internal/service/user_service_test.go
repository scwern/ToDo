package service

import (
	"ToDo/internal/domain/user"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByEmail(email string) (*user.User, error) {
	args := m.Called(email)
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserRepository) GetAll() ([]user.User, error) {
	args := m.Called()
	return args.Get(0).([]user.User), args.Error(1)
}

func (m *MockUserRepository) GetById(id uuid.UUID) (*user.User, error) {
	args := m.Called(id)
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserRepository) Create(u user.User) (user.User, error) {
	args := m.Called(u)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockUserRepository) Update(id uuid.UUID, updated user.User) (*user.User, error) {
	args := m.Called(id, updated)
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func createTestUser() user.User {
	u := user.NewUser("Test", "test@test.com", "password")
	u.SetID(uuid.New())
	return u
}

func TestUserService_GetByEmail(t *testing.T) {
	repo := new(MockUserRepository)
	service := NewUserService(repo)

	expectedUser := createTestUser()
	repo.On("GetByEmail", expectedUser.Email()).Return(&expectedUser, nil)

	result, err := service.GetByEmail(expectedUser.Email())
	assert.NoError(t, err)
	assert.Equal(t, &expectedUser, result)
}

func TestUserService_GetByEmail_Error(t *testing.T) {
	repo := new(MockUserRepository)
	service := NewUserService(repo)
	email := "error@test.com"
	repo.On("GetByEmail", email).Return((*user.User)(nil), errors.New("error"))
	_, err := service.GetByEmail(email)
	assert.Error(t, err)
}

func TestUserService_GetAll(t *testing.T) {
	repo := new(MockUserRepository)
	service := NewUserService(repo)

	user1 := createTestUser()
	user2 := createTestUser()
	expectedUsers := []user.User{user1, user2}

	repo.On("GetAll").Return(expectedUsers, nil)
	users, err := service.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestUserService_GetById(t *testing.T) {
	repo := new(MockUserRepository)
	service := NewUserService(repo)

	expectedUser := createTestUser()
	repo.On("GetById", expectedUser.ID()).Return(&expectedUser, nil)

	result, err := service.GetById(expectedUser.ID())
	assert.NoError(t, err)
	assert.Equal(t, &expectedUser, result)
}

func TestUserService_Create(t *testing.T) {
	repo := new(MockUserRepository)
	service := NewUserService(repo)

	newUser := user.NewUser("New", "new@test.com", "pass")
	createdUser := newUser
	createdUser.SetID(uuid.New())

	repo.On("Create", newUser).Return(createdUser, nil)
	result, err := service.Create(newUser)
	assert.NoError(t, err)
	assert.Equal(t, createdUser, result)
}

func TestUserService_Create_Error(t *testing.T) {
	repo := new(MockUserRepository)
	service := NewUserService(repo)

	newUser := user.NewUser("Error", "error@test.com", "pass")
	repo.On("Create", newUser).Return(user.User{}, errors.New("error"))

	_, err := service.Create(newUser)
	assert.Error(t, err)
}

func TestUserService_Update(t *testing.T) {
	repo := new(MockUserRepository)
	service := NewUserService(repo)

	userID := uuid.New()
	updatedUser := createTestUser()
	updatedUser.SetID(userID)
	updatedUser.SetName("Updated")

	repo.On("Update", userID, updatedUser).Return(&updatedUser, nil)
	result, err := service.Update(userID, updatedUser)
	assert.NoError(t, err)
	assert.Equal(t, &updatedUser, result)
}
