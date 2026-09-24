package lessonusecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/apperr"
	"github.com/mrruke12/lms/internal/domains/auth"
	"github.com/mrruke12/lms/internal/domains/lesson"
)

type CreateLessonCommand struct {
	actor auth.Actor
	name  string
}

func (s *Service) CreateLesson(ctx context.Context, cmd CreateLessonCommand) (*uuid.UUID, error) {
	if !cmd.actor.HasRole(auth.RoleTeacher) || !cmd.actor.HasPermission(auth.PermissionLessonCreate) {
		return nil, apperr.Forbidden()
	}

	lesson, err := lesson.NewLesson(
		cmd.actor.UserID(),
		cmd.name,
	)

	if err != nil {
		return nil, err // TODO: fix errors handling after fixing the apperr package itself
	}

	id, err := s.lessons.Create(ctx, lesson)

	if err != nil {
		return nil, err
	}

	return id, nil
}
