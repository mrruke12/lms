package group

import "github.com/mrruke12/lms/pkg/enum"

type Degree string

var (
	DegreeAssociate Degree = "associate"
	DegreeBachelor  Degree = "bachelor"
	DegreeMaster    Degree = "master"
	DegreeDoctor    Degree = "doctor"
)

var degreeSet = enum.NewSet(
	DegreeAssociate,
	DegreeBachelor,
	DegreeMaster,
	DegreeDoctor,
)
