package auth

import "github.com/mrruke12/lms/pkg/enum"

type Permission string

const (
	PermissionLessonCreate Permission = "lesson:create"
	PermissionLessonUpdate Permission = "lesson:update"
	PermissionLessonDelete Permission = "lesson:delete"
	PermissionLessonRead   Permission = "lesson:read"
)

var permissionSet = enum.NewSet(
	PermissionLessonCreate,
	PermissionLessonUpdate,
	PermissionLessonDelete,
	PermissionLessonRead,
)
