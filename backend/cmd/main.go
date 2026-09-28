package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/application/lessonusecase"
	"github.com/mrruke12/lms/internal/domains/assessment"
	"github.com/mrruke12/lms/internal/domains/auth"
	"github.com/mrruke12/lms/internal/domains/element"
	"github.com/mrruke12/lms/internal/infrastructure/config"
	"github.com/mrruke12/lms/internal/infrastructure/postgres"
)

func main() {
	config.LoadEnvironment()

	cfg, _ := config.GetDBConfig()

	// postgres.RunMigrations(*cfg)

	pool := postgres.Connect(context.Background(), *cfg)

	lessonRepo := postgres.NewLessonRepository(pool)
	elementRepo := postgres.NewElementRepository(pool)
	revisionRepo := postgres.NewRevisionRepository(pool)

	serv := lessonusecase.NewService(
		lessonRepo,
		elementRepo,
		revisionRepo,
	)

	ctx := context.Background()
	userID := uuid.MustParse("01a0e79a-9445-7495-98a5-e5fc5d4ff672")
	actor := auth.NewActor(
		userID, auth.RoleTeacher, auth.PermissionSet.Values(),
	)

	l, err := serv.CreateLesson(
		ctx,
		lessonusecase.CreateLessonCommand{
			Actor: *actor,
			Name:  "test lesson 111111",
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(l.ID())

	err = serv.PublishLesson(ctx, lessonusecase.PublishLessonCommand{
		Actor:    *actor,
		LessonID: l.ID(),
	})

	if err != nil {
		panic(err)
	}

	els := []element.Element{
		el(l.ID(), 1),
		el(l.ID(), 2),
		el(l.ID(), 3),
	}

	err = serv.UpdateLesson(ctx, lessonusecase.UpdateLessonCommand{
		Actor:    *actor,
		LessonID: l.ID(),
		Elements: els,
	})

	if err == nil {
		panic("expected error")
	}

	fmt.Println(err)

	err = serv.UnpublishLesson(ctx, lessonusecase.UnpublishLessonCommand{
		Actor:    *actor,
		LessonID: l.ID(),
	})

	if err != nil {
		panic(err)
	}

	err = serv.UpdateLesson(ctx, lessonusecase.UpdateLessonCommand{
		Actor:    *actor,
		LessonID: l.ID(),
		Elements: els,
	})

	if err != nil {
		panic(err)
	}

	err = els[0].ChangeConfig(json.RawMessage(`[{"test": "123"}]`))

	if err != nil {
		panic(err)
	}

	els2 := []element.Element{
		els[0],
		els[1],
	}

	err = serv.UpdateLesson(ctx, lessonusecase.UpdateLessonCommand{
		Actor:    *actor,
		LessonID: l.ID(),
		Elements: els2,
	})
}

func el(id uuid.UUID, num int) element.Element {
	elm, err := element.NewElement(
		id,
		nil,
		element.TypeTest,
		assessment.AssessmentNone,
		json.RawMessage(fmt.Sprintf(`[{"test": "%d"}]`, num)),
	)

	if err != nil {
		panic(err)
	}

	return *elm
}
