package element

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Element, error)
	GetByLessonID(ctx context.Context, id uuid.UUID) ([]Element, error)
	BulkCreate(ctx context.Context, elements []Element) error
	BulkUpdate(ctx context.Context, elements []Element) error
	BulkDelete(ctx context.Context, ids []uuid.UUID) error
}
