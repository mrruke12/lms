package lessonusecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/application/apperr"
	"github.com/mrruke12/lms/internal/domains/auth"
	"github.com/mrruke12/lms/internal/domains/lesson"
)

type UnpublishLessonCommand struct {
	Actor    auth.Actor
	LessonID uuid.UUID
}

func (s *Service) UnpublishLesson(ctx context.Context, cmd UnpublishLessonCommand) error {
	if !cmd.Actor.HasRole(auth.RoleTeacher) || !cmd.Actor.HasPermission(auth.PermissionLessonUpdate) {
		return apperr.Error(apperr.CodePermissionDenied, "has no permission to unpublish lesson")
	}

	l, err := s.lessons.GetByID(ctx, cmd.LessonID)

	if err != nil {
		return apperr.MapDBErr(err)
	}

	err = l.SetStatus(lesson.StatusEditing)

	if err != nil {
		return apperr.Wrap(apperr.CodePermissionDenied, "cannot set status", err)
	}

	err = s.lessons.Update(ctx, l)

	if err != nil {
		return apperr.MapDBErr(err)
	}

	return nil
}
