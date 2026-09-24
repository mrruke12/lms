package lessonusecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/application/apperr"
	"github.com/mrruke12/lms/internal/domains/auth"
	"github.com/mrruke12/lms/internal/domains/lesson"
)

type CreateLessonCommand struct {
	actor auth.Actor
	name  string
}

func (s *Service) CreateLesson(ctx context.Context, cmd CreateLessonCommand) (*uuid.UUID, error) {
	if !cmd.actor.HasRole(auth.RoleTeacher) || !cmd.actor.HasPermission(auth.PermissionLessonCreate) {
		return nil, apperr.Error("access_denied", "access denied")
	}

	lesson, err := lesson.NewLesson(
		cmd.actor.UserID(),
		cmd.name,
	)

	if err != nil {
		return nil, apperr.Wrap("invalid_lesson", "failed to create lesson", err)
	}

	id, err := s.lessons.Create(ctx, lesson)

	if err != nil {
		return nil, apperr.Wrap("save_failed", "failed to save new lesson", err)
	}

	return id, nil
}
