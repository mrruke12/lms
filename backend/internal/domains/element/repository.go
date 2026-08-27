package element

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Element, error)
	GetAllByLessonID(ctx context.Context, id uuid.UUID) ([]Element, error)
	GetTypes(ctx context.Context) ([]Type, error)
}
