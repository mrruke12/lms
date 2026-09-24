package auth

import (
	"github.com/google/uuid"
)

type Actor struct {
	userID      uuid.UUID
	role        Role
	permissions map[Permission]struct{}
}

func NewActor(userID uuid.UUID, role Role, permissions []Permission) *Actor {
	actor := &Actor{
		userID: userID,
		role:   role,
	}

	actor.permissions = make(map[Permission]struct{})

	for _, permission := range permissions {
		actor.permissions[permission] = struct{}{}
	}

	return actor
}

func (a *Actor) Role() Role {
	return a.role
}

func (a *Actor) HasRole(role Role) bool {
	return a.role == role
}

func (a *Actor) HasPermission(permission Permission) bool {
	_, exist := a.permissions[permission]

	return exist
}

func (a *Actor) UserID() uuid.UUID {
	return a.userID
}
