package enrollment

import "github.com/mrruke12/lms/pkg/enum"

type Status string

var (
	EnrollmentActive  Status = "active"
	EnrollmentArchive Status = "archive"
)

var enrollmentStatusSet = enum.NewSet(
	EnrollmentActive,
	EnrollmentArchive,
)

var enrollmentStatusTransitions = enum.NewStateMachine(
	enrollmentStatusSet,
	map[Status][]Status{
		EnrollmentActive: {
			EnrollmentArchive,
		},
	},
)
