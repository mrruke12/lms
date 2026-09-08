package assessment

import "github.com/mrruke12/lms/pkg/enum"

type Type string

var (
	AssessmentNone       Type = "none"
	AssessmentManual     Type = "manual"
	AssessmentPredefined Type = "predefined"
	AssessmentTests      Type = "tests"
)

var assessmentSet = enum.NewSet(
	AssessmentNone,
	AssessmentManual,
	AssessmentPredefined,
	AssessmentTests,
)

var IsValidType = assessmentSet.Has
