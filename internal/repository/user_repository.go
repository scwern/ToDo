package repository

import (
	"ToDo/internal/domain/user"
	"errors"
)

type UserRepository struct {
	users  []user.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make([]user.User, 0),
		nextID: 1,
	}
}

func (r *UserRepository) GetAll() []user.User {
	return r.users
}

func (r *UserRepository) GetById(id int) (*user.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Create(u user.User) user.User {
	u.ID = r.nextID
	r.nextID++
	r.users = append(r.users, u)
	return u
}

func (r *UserRepository) Update(id int, updated user.User) (*user.User, error) {
	for i, u := range r.users {
		if u.ID == id {
			updated.ID = id
			r.users[i] = updated
			return &r.users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Delete(id int) error {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
