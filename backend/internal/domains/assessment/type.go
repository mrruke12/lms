package assessment

import "github.com/mrruke12/lms/pkg/enum"

type Type string

var (
	None       Type = "none"
	Manual     Type = "manual"
	Predefined Type = "predefined"
	Tests      Type = "tests"
)

var assessmentSet = enum.NewSet(
	None,
	Manual,
	Predefined,
	Tests,
)

var IsValidType = assessmentSet.Has
