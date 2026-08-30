package submission

import "github.com/mrruke12/lms/pkg/enum"

type Status string

var (
	StatusPending   Status = "pending"
	StatusEvaluated Status = "evaluated"
)

var statusSet = enum.NewSet(
	StatusPending,
	StatusEvaluated,
)

var statusTransitions = enum.NewStateMachine(
	statusSet,
	map[Status][]Status{
		StatusPending: {
			StatusEvaluated,
		},
	},
)

var IsValidStatus = statusSet.Has
