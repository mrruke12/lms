package lesson

import "github.com/mrruke12/lms/pkg/enum"

type Status string

var (
	StatusDraft     Status = "draft"
	StatusEditing   Status = "editing"
	StatusPublished Status = "published"
	StatusArchived  Status = "archived"
)

var statusSet = enum.NewSet(
	StatusDraft,
	StatusEditing,
	StatusPublished,
	StatusArchived,
)

var statusTransitions = enum.NewStateMachine(
	statusSet,
	map[Status][]Status{
		StatusDraft: {
			StatusPublished,
		},
		StatusPublished: {
			StatusEditing,
		},
		StatusEditing: {
			StatusPublished,
			StatusArchived,
		},
		StatusArchived: {
			StatusDraft,
		},
	},
)

var IsValidStatus = statusSet.Has
