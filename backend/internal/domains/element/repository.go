package element

import (
	"context"

	"github.com/google/uuid"
)

type ElementRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Element, error)
	GetAllByLessonID(ctx context.Context, id uuid.UUID) ([]Element, error)
}

type TypeRepository interface {
	GetByID(ctx context.Context, id int) (*Type, error)
}
