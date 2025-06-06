package in_memory

import (
	"ToDo/internal/domain/user"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type UserRepository struct {
	users map[uuid.UUID]user.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[uuid.UUID]user.User),
	}
}

func (r *UserRepository) GetByEmail(email string) (*user.User, error) {
	for _, u := range r.users {
		if u.Email() == email {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) GetAll() ([]user.User, error) {
	result := make([]user.User, 0, len(r.users))
	for _, u := range r.users {
		result = append(result, u)
	}
	return result, nil
}

func (r *UserRepository) GetById(id uuid.UUID) (*user.User, error) {
	if u, exists := r.users[id]; exists {
		return &u, nil
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Create(u user.User) (user.User, error) {
	fmt.Println("Adding user to repository:", u)
	if u.ID() == uuid.Nil {
		u.SetID(uuid.New())
	}
	r.users[u.ID()] = u
	return u, nil
}

func (r *UserRepository) Update(id uuid.UUID, updated user.User) (*user.User, error) {
	if existingUser, exists := r.users[id]; exists {
		existingUser.SetName(updated.Name())
		existingUser.SetEmail(updated.Email())
		existingUser.SetPassword(updated.Password())

		r.users[id] = existingUser

		return &existingUser, nil
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	if _, exists := r.users[id]; exists {
		delete(r.users, id)
		return nil
	}
	return errors.New("user not found")
}
