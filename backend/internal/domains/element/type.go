package element

import "github.com/mrruke12/lms/pkg/enum"

// Element type defined in DB
type Type string

const (
	TypeTest Type = "test"
)

var typeSet = enum.NewSet[Type](TypeTest)
