package lesson

type LessonStatus string

type lessonStatusEnum struct {
	Draft     LessonStatus
	Editing   LessonStatus
	Published LessonStatus
	Archived  LessonStatus
}

var Status = lessonStatusEnum{
	"draft",
	"editing",
	"published",
	"archived",
}
