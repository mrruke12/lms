package user

import "github.com/mrruke12/lms/pkg/enum"

type Type string

var (
	UserStudent Type = "student"
	UserTeacher Type = "teacher"
)

var userTypeSet = enum.NewSet(
	UserStudent,
	UserTeacher,
)
