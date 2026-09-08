package user

import "github.com/google/uuid"

// Base user entity
type User struct {
	id uuid.UUID

	typ Type

	login        string
	passwordHash string

	name       string
	surname    string
	patronymic string
}

/*
Getters
*/

func (u *User) Type() Type {
	return u.typ
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Login() string {
	return u.login
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Surname() string {
	return u.surname
}

func (u *User) Patronymic() string {
	return u.patronymic
}
