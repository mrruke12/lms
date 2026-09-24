package lessonusecase

import "github.com/mrruke12/lms/internal/domains/lesson"

type Service struct {
	lessons lesson.Repository
}
