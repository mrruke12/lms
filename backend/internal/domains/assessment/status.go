package assessment

import "github.com/mrruke12/lms/pkg/enum"

type Status string

var (
	StatusPending   Status = "pending"
	StatusPartial   Status = "partial"
	StatusEvaluated Status = "evaluated"
)

var statusSet = enum.NewSet(
	StatusPending,
	StatusPartial,
	StatusEvaluated,
)

var statusTransitions = enum.NewStateMachine(
	statusSet,
	map[Status][]Status{
		StatusPending: {
			StatusPartial,
			StatusEvaluated,
		},
		StatusPartial: {
			StatusEvaluated,
		},
	},
)

var IsValidStatus = statusSet.Has
