package group

import (
	"github.com/google/uuid"
)

type Group struct {
	id             uuid.UUID
	name           string
	degree         Degree
	entryYear      int
	graduationYear int
}
