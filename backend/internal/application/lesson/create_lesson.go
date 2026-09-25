package lessonusecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/application/apperr"
	"github.com/mrruke12/lms/internal/domains/auth"
	"github.com/mrruke12/lms/internal/domains/lesson"
)

type CreateLessonCommand struct {
	Actor auth.Actor
	Name  string
}

func (s *Service) CreateLesson(ctx context.Context, cmd CreateLessonCommand) (*uuid.UUID, error) {
	if !cmd.Actor.HasRole(auth.RoleTeacher) || !cmd.Actor.HasPermission(auth.PermissionLessonCreate) {
		return nil, apperr.Error(apperr.CodePermissionDenied, "cannot create lessons")
	}

	l, err := lesson.NewLesson(
		cmd.Actor.UserID(),
		cmd.Name,
	)

	if err != nil {
		return nil, apperr.Wrap("invalid_lesson", "failed to create lesson", err)
	}

	id, err := s.lessons.Create(ctx, l)

	if err != nil {
		return nil, apperr.Wrap("save_failed", "failed to save new lesson", err)
	}

	return id, nil
}
