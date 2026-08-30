package attempt

import "github.com/mrruke12/lms/pkg/enum"

type Status string

var (
	StatusActive    Status = "active"
	StatusSubmitted Status = "submitted"
	StatusCompleted Status = "completed"
	StatusOverdue   Status = "overdue"
	StatusArchived  Status = "archived"
)

var statusSet = enum.NewSet(
	StatusActive,
	StatusSubmitted,
	StatusCompleted,
	StatusOverdue,
	StatusArchived,
)

var statusTransitions = enum.NewStateMachine(
	statusSet,
	map[Status][]Status{
		StatusActive: {
			StatusSubmitted,
			StatusOverdue,
			StatusArchived,
		},
		StatusSubmitted: {
			StatusCompleted,
			StatusArchived,
		},
		StatusCompleted: {
			StatusArchived,
		},
	},
)

func IsValidStatus(status Status) bool {
	return statusSet.Has(status)
}
