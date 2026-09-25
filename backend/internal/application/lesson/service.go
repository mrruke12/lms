package lessonusecase

import (
	"github.com/mrruke12/lms/internal/domains/element"
	"github.com/mrruke12/lms/internal/domains/lesson"
	"github.com/mrruke12/lms/internal/domains/revision"
)

type Service struct {
	lessons   lesson.Repository
	elements  element.Repository
	revisions revision.Repository
}
