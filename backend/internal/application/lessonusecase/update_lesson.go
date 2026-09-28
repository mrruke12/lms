package lessonusecase

import (
	"context"
	"slices"

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

	persistedEls, newEls, deletedEls := element.Categorize(cmd.Elements, els)

	err = s.deleteElements(ctx, deletedEls)

	if err != nil {
		return err
	}

	if len(newEls) > 0 {
		err = s.elements.BulkCreate(ctx, newEls)
	}

	if err != nil {
		return err
	}

	return s.updateElements(ctx, slices.Concat(persistedEls, newEls), revs)
}

func (s *Service) deleteElements(ctx context.Context, els []element.Element) error {
	if len(els) == 0 {
		return nil
	}

	ids := make([]uuid.UUID, len(els))

	for i := range els {
		ids[i] = els[i].ID()
	}

	err := s.elements.BulkDelete(ctx, ids)

	return err
}

func (s *Service) updateElements(ctx context.Context, els []element.Element, revs []revision.Revision) error {
	if len(els) == 0 {
		return nil
	}

	updatedEls, newRevs := revision.ComputeRevisions(els, revs)

	if len(newRevs) > 0 {
		err := s.revisions.BulkCreate(ctx, newRevs)

		if err != nil {
			return apperr.MapDBErr(err)
		}
	}

	if len(updatedEls) > 0 {
		return s.elements.BulkUpdate(ctx, updatedEls)
	}

	return nil
}
