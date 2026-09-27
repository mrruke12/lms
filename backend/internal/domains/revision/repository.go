package revision

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetByLessonID(ctx context.Context, id uuid.UUID) ([]Revision, error)
	BulkCreate(ctx context.Context, revisions []Revision) error
}
