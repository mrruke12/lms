package element

type AssessmentType string

type assessmentTypes struct {
	None       AssessmentType
	Manual     AssessmentType
	Predefined AssessmentType
	Tests      AssessmentType
}

var Assessment = assessmentTypes{
	None:       "none",
	Manual:     "manual",
	Predefined: "predefined",
	Tests:      "tests",
}
