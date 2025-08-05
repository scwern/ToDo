package service

import (
	"ToDo/internal/domain/user"
	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	GetByEmail(email string) (*user.User, error)
	GetAll() ([]user.User, error)
	GetById(id uuid.UUID) (*user.User, error)
	Create(u user.User) (user.User, error)
	Update(id uuid.UUID, updated user.User) (*user.User, error)
	Delete(id uuid.UUID) error
}

type UserService struct {
	repo UserRepositoryInterface
}

// NewUserService creates a new user service instance
// @Summary Create user service
// @Tags Users
// @Router /user-service [post]
func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

// GetByEmail finds user by email
// @Summary Find user by email
// @Tags Users
// @Param email query string true "User email"
// @Success 200 {object} user.User
// @Router /users/search [get]
func (s *UserService) GetByEmail(email string) (*user.User, error) {
	return s.repo.GetByEmail(email)
}

// GetAll returns all users
// @Summary Get all users
// @Tags Users
// @Success 200 {array} user.User
// @Router /users [get]
func (s *UserService) GetAll() ([]user.User, error) {
	return s.repo.GetAll()
}

// GetById returns user by ID
// @Summary Get user by ID
// @Tags Users
// @Param id path string true "User UUID"
// @Success 200 {object} user.User
// @Router /users/{id} [get]
func (s *UserService) GetById(id uuid.UUID) (*user.User, error) {
	return s.repo.GetById(id)
}

// Create creates new user
// @Summary Create user
// @Tags Users
// @Param input body user.User true "User data"
// @Success 200 {object} user.User
// @Router /users [post]
func (s *UserService) Create(u user.User) (user.User, error) {
	return s.repo.Create(u)
}

// Update updates existing user
// @Summary Update user
// @Tags Users
// @Param id path string true "User UUID"
// @Param input body user.User true "Updated user data"
// @Success 200 {object} user.User
// @Router /users/{id} [put]
func (s *UserService) Update(id uuid.UUID, updated user.User) (*user.User, error) {
	return s.repo.Update(id, updated)
}

// Delete removes user
// @Summary Delete user
// @Tags Users
// @Param id path string true "User UUID"
// @Success 200
// @Router /users/{id} [delete]
func (s *UserService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
