package enrollment

import "github.com/mrruke12/lms/pkg/enum"

type Type string

var (
	EnrollmentCourse Type = "course"
	EnrollmentGroup  Type = "group"
)

var enrollmentTypeSet = enum.NewSet(
	EnrollmentCourse,
	EnrollmentGroup,
)
