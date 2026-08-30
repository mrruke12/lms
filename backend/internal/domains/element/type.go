package element

import (
	"time"

	"github.com/google/uuid"
)

// Element type defined in DB
type Type struct {
	ID  int64
	Key string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	DeletedBy *uuid.UUID
}
