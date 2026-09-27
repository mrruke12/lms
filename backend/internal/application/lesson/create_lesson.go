package lessonusecase

import (
	"context"

	"github.com/mrruke12/lms/internal/application/apperr"
	"github.com/mrruke12/lms/internal/domains/auth"
	"github.com/mrruke12/lms/internal/domains/lesson"
)

type CreateLessonCommand struct {
	Actor auth.Actor
	Name  string
}

func (s *Service) CreateLesson(ctx context.Context, cmd CreateLessonCommand) error {
	if !cmd.Actor.HasRole(auth.RoleTeacher) || !cmd.Actor.HasPermission(auth.PermissionLessonCreate) {
		return apperr.Error(apperr.CodePermissionDenied, "has no permission to create lessons")
	}

	l, err := lesson.NewLesson(
		cmd.Actor.UserID(),
		cmd.Name,
	)

	if err != nil {
		return apperr.Wrap("invalid_lesson", "failed to create lesson", err)
	}

	err = s.lessons.Create(ctx, l)

	if err != nil {
		return apperr.Wrap("save_failed", "failed to save new lesson", err)
	}

	return nil
}
