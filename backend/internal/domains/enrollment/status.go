package enrollment

import "github.com/mrruke12/lms/pkg/enum"

type Status string

var (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
)

var enrollmentStatusSet = enum.NewSet(
	StatusActive,
	StatusArchived,
)

var enrollmentStatusTransitions = enum.NewStateMachine(
	enrollmentStatusSet,
	map[Status][]Status{
		StatusActive: {
			StatusArchived,
		},
	},
)
