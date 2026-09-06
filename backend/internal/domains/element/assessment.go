package element

import "github.com/mrruke12/lms/pkg/enum"

type AssessmentType string

var (
	None       AssessmentType = "none"
	Manual     AssessmentType = "manual"
	Predefined AssessmentType = "predefined"
	Tests      AssessmentType = "tests"
)

var assessmentSet = enum.NewSet(
	None,
	Manual,
	Predefined,
	Tests,
)

var IsValidAssessmentType = assessmentSet.Has
