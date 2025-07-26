package user

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	name     string
	email    string
	password string
}

func NewUser(name, email, password string) User {
	return User{
		id:       uuid.New(),
		name:     name,
		email:    email,
		password: password,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetID(id uuid.UUID) {
	u.id = id
}
