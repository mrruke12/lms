package auth

import "github.com/mrruke12/lms/pkg/enum"

type Role string

const (
	RoleStudent Role = "student"
	RoleTeacher Role = "Teacher"
)

var roleSet = enum.NewSet(
	RoleStudent,
	RoleTeacher,
)
