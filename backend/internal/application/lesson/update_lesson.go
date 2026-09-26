package lessonusecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/application/apperr"
	"github.com/mrruke12/lms/internal/domains/auth"
	"github.com/mrruke12/lms/internal/domains/element"
	"github.com/mrruke12/lms/internal/domains/revision"
)

type UpdateLessonCommand struct {
	Actor    auth.Actor
	LessonID uuid.UUID
	Elements []element.Element
}

func (s *Service) UpdateLesson(ctx context.Context, cmd UpdateLessonCommand) error {
	if !cmd.Actor.HasRole(auth.RoleTeacher) || !cmd.Actor.HasPermission(auth.PermissionLessonUpdate) {
		return apperr.Error(apperr.CodePermissionDenied, "has no permission to edit lessons")
	}

	l, err := s.lessons.GetByID(ctx, cmd.LessonID)

	if err != nil {
		return apperr.MapDBErr(err)
	}

	if !l.CanEdit() {
		return apperr.Error(apperr.CodePermissionDenied, "cannot edit lesson with status "+string(l.Status()))
	}

	els, err := s.elements.GetByLessonID(ctx, cmd.LessonID)

	if err != nil {
		return apperr.MapDBErr(err)
	}

	revs, err := s.revisions.GetByLessonID(ctx, cmd.LessonID)

	if err != nil {
		return apperr.MapDBErr(err)
	}

	newRevs := revision.ComputeRevisions(els, revs)

	if len(newRevs) > 0 {
		err := s.revisions.BulkCreate(ctx, newRevs)

		if err != nil {
			return apperr.MapDBErr(err)
		}
	}

	return nil
}
